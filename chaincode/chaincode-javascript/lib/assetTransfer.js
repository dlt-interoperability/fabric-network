/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

"use strict";

const { Contract } = require("fabric-contract-api");

class AssetTransfer extends Contract {
  async InitLedger(ctx) {
    const key = "key0";
    const value = "value0";
    await ctx.stub.putState(key, Buffer.from(value));
    console.info(`Ledger initialized with key ${key}`);
  }

  // CreateAsset issues a new asset to the world state with given details.
  async CreateAsset(ctx, key, value) {
    return ctx.stub.putState(key, Buffer.from(value));
  }

  // ReadAsset returns the asset stored in the world state with given key.
  async ReadAsset(ctx, key) {
    const value = await ctx.stub.getState(key); // get the asset from chaincode state
    if (!value) {
      throw new Error(`The asset ${key} does not exist`);
    }
    return value;
  }

  // UpdateAsset updates an existing asset in the world state with provided parameters.
  async UpdateAsset(ctx, key, value) {
    const exists = await this.AssetExists(ctx, key);
    if (!exists) {
      throw new Error(`The asset ${key} does not exist`);
    }

    // overwriting original asset with new asset
    return ctx.stub.putState(key, Buffer.from(value));
  }

  // DeleteAsset deletes an given asset from the world state.
  async DeleteAsset(ctx, key) {
    const exists = await this.AssetExists(ctx, key);
    if (!exists) {
      throw new Error(`The asset ${key} does not exist`);
    }
    return ctx.stub.deleteState(key);
  }

  // AssetExists returns true when asset with given ID exists in world state.
  async AssetExists(ctx, key) {
    const value = await ctx.stub.getState(key);
    return value && value.length > 0;
  }

  // GetAllAssets returns all assets found in the world state.
  async GetAllAssets(ctx) {
    const allResults = [];
    // range query with empty string for startKey and endKey does an open-ended query of all assets in the chaincode namespace.
    const iterator = await ctx.stub.getStateByRange("", "");
    let result = await iterator.next();
    while (!result.done) {
      const strValue = Buffer.from(result.value.value.toString()).toString(
        "utf8",
      );
      allResults.push({ Key: result.value.key, Value: strValue });
      result = await iterator.next();
    }
    return JSON.stringify(allResults);
  }

  async GetHistoryForKey(ctx, key) {
    console.log(`Get history for key ${key}`);
    let iterator = await ctx.stub.getHistoryForKey(key);
    let allResults = [];
    let res = await iterator.next();
    while (!res.done) {
      if (res.value && res.value.value.toString()) {
        let jsonRes = {};
        console.log(res.value.value.toString("utf8"));
        jsonRes.txId = res.value.tx_id;
        jsonRes.timestamp = res.value.timestamp;
        jsonRes.isDelete = res.value.isDelete;
        try {
          jsonRes.value = JSON.parse(res.value.value.toString("utf8"));
        } catch (err) {
          console.log(err);
          jsonRes.value = res.value.value.toString("utf8");
        }
        allResults.push(jsonRes);
      }
      res = await iterator.next();
    }
    iterator.close();
    return JSON.stringify(allResults);
  }
}

module.exports = AssetTransfer;
