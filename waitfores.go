// Copyright © 2019 blacktop
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package waitfores

import (
	"context"
	"time"

	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// WaitForEs contains the wait-for-es settings
type WaitForEs struct {
	URL      string
	Username string
	Password string
	Health   string
	Timeout  int64
}

func (wfe *WaitForEs) testConnection() error {

	client, err := elastic.NewSimpleClient(
		elastic.SetURL(wfe.URL),
		elastic.SetBasicAuth(
			wfe.Username,
			wfe.Password,
		),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create elasticsearch simple client")
	}

	// Ping the Elasticsearch server to get e.g. the version number
	log.Debugf("attempting to PING to: %s", wfe.URL)
	info, code, err := client.Ping(wfe.URL).Do(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to ping elasticsearch")
	}

	log.WithFields(log.Fields{
		"code":    code,
		"cluster": info.ClusterName,
		"version": info.Version.Number,
		"url":     wfe.URL,
	}).Debug("elasticSearch connection successful")

	return nil
}

// WaitForConnection waits for connection to Elasticsearch to be ready
func (wfe *WaitForEs) WaitForConnection(ctx context.Context, timeout int64) error {

	var err error

	secondsWaited := 0

	connCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	log.Info("===> trying to connect to elasticsearch")
	for {
		// Try to connect to Elasticsearch
		select {
		case <-connCtx.Done():
			return errors.Wrapf(err, "connecting to elasticsearch timed out after %d seconds", secondsWaited)
		default:
			err = wfe.testConnection()
			if err == nil {
				log.Infof("elasticsearch came online after %d seconds", secondsWaited)
				return nil
			}
			// not ready yet
			secondsWaited++
			log.Debug(" * could not connect to elasticsearch (sleeping for 1 second)")
			time.Sleep(1 * time.Second)
		}
	}
}
