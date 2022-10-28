package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/best-bank-api/util"
)

func createRandomEntries(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry

}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntries(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entryMock := createRandomEntries(t, account)

	entry2, err := testQueries.GetEntry(context.Background(), entryMock.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entryMock.ID, entry2.ID)
	require.Equal(t, entryMock.AccountID, entry2.AccountID)
	require.Equal(t, entryMock.Amount, entry2.Amount)
	require.WithinDuration(t, entryMock.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntries(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Offset:    5,
		Limit:     5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}

}
