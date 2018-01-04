// Copyright 2017 Orinoco Payments
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package etherutils

// Watch for a transaction to be mined.
// If the transaction is mined this will return true
// func Mined(client *ethclient.Client, tx *types.Transaction) (mined bool, err error) {
// 	type ContextKey string
// 	key := ContextKey("hash")
//
// 	headChan := make(chan *types.Header)
// 	// d := time.Now().Add(60 * time.Second)
// 	// ctx, cancel := context.WithDeadline(context.Background(), d)
// 	// defer cancel()
// 	ctx := context.WithValue(context.Background(), key, tx.Hash().Hex())
//
// 	_, err = client.SubscribeNewHead(ctx, headChan)
// 	if err != nil {
// 		return false, errors.New("Failed to listen for confirmation")
// 	}
// 	pending := true
// 	for pending {
// 		_ = <-headChan
// 		_, pending, err = client.TransactionByHash(ctx, tx.Hash())
// 		if err != nil {
// 			// An error of some sort occurred
// 			return false, err
// 		}
// 	}
//
// 	return true, nil
// }
