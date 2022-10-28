package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/best-bank-api/util"
)

func createRandomTransfer(t *testing.T, fromAccount Account, toAccount Account) Transfer {

	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	toAccount := createRandomAccount(t)
	fromAccount := createRandomAccount(t)
	createRandomTransfer(t, fromAccount, toAccount)
}

func TestGetTransfer(t *testing.T) {
	toAccount := createRandomAccount(t)
	fromAccount := createRandomAccount(t)
	transferMock := createRandomTransfer(t, fromAccount, toAccount)

	transfer, err := testQueries.GetTransfer(context.Background(), transferMock.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transferMock.ID, transfer.ID)
	require.Equal(t, transferMock.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transferMock.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transferMock.Amount, transfer.Amount)
	require.WithinDuration(t, transferMock.CreatedAt, transfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	arg := ListTransfersParams{
		ToAccountID:   account1.ID,
		FromAccountID: account1.ID,
		Offset:        5,
		Limit:         5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}

}
