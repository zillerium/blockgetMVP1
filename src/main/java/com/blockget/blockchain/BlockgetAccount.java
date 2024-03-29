package com.blockget.blockchain;

import ch.decent.sdk.DCoreApi;
import ch.decent.sdk.crypto.Address;
import ch.decent.sdk.crypto.Credentials;
import ch.decent.sdk.model.Account;
import ch.decent.sdk.model.AssetAmount;
import ch.decent.sdk.model.Fee;
import ch.decent.sdk.model.TransactionConfirmation;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;


@Component
public class BlockgetAccount {
    // connection to accounts
    private static final Long AMOUNT_OF_DCT_REQUIRED_FOR_CREATION = 100000L;


    @Autowired
    private BlockgetLogin loginExample;
    @Autowired
    private BlockgetGenerateKeys generateKeys;

    /**
     * Example of getting any account by its name.
     *
     * @param accountName name of the account e.g. dw-account
     * @return Account instance for given account name
     */
    public Account getAccountByName(String accountName) {

        BlockgetConnection connectionExample = new BlockgetConnection();
        final DCoreApi dcoreApi = connectionExample.connect();

        return dcoreApi
                .getAccountApi()
                .getByName(accountName)
                .blockingGet();
    }

    /**
     * Example of account creation with initial fee.
     *
     * @param newAccountName Unique account name that you wish to create.
     * @return Confirmation about transaction
     */
    public TransactionConfirmation createAccount(String newAccountName) {

        BlockgetConnection connectionExample = new BlockgetConnection();
        final DCoreApi dcoreApi = connectionExample.connect();
        final Credentials credentials = loginExample.login(connectionExample);
        final Address newAccountPublicKey = generateKeys.generateKeys();
        final AssetAmount dctAssetAmount = new AssetAmount(AMOUNT_OF_DCT_REQUIRED_FOR_CREATION);
        final Fee initialFee = new Fee(dctAssetAmount.getAssetId(), AMOUNT_OF_DCT_REQUIRED_FOR_CREATION);

        return dcoreApi.getAccountApi().create(
                credentials,
                newAccountName,
                newAccountPublicKey,
                initialFee
        ).blockingGet();
    }
}
