package simpleasset_test

import (
	"testing"

	sa "simplecontract/simpleasset"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
)

func TestSimpleAssetFunction(t *testing.T) {
	// assert equality
	assert.Equal(t, 1, 1, "they should be equal")
}

func TestInit(t *testing.T) {

	// Arrange
	sacc := new(sa.SimpleAsset)                   // chaincode instance
	stub := shim.NewMockStub("simpleasset", sacc) // shim stub

	// 傳入 0 個參數，Init應失敗
	t.Run("Init with 0 argument should fail", func(t *testing.T) {

		// Act
		res := stub.MockInit("1", nil)

		// Assert
		assert.Equal(t, int(res.Status), shim.ERROR)
		assert.Equal(t, res.Message, "Incorrect arguments. Expecting a key and a value")
	})

	// 傳入 2 個參數，Init應成功
	t.Run("Init with 2 argument should success", func(t *testing.T) {

		// Arrange
		args := [][]byte{
			[]byte("a"),
			[]byte("12345"),
		}

		// Act
		res := stub.MockInit("1", args)

		// Assert
		assert.Equal(t, int(res.Status), shim.OK)
	})
}

// Init 之後，應能立即查詢到 key-value 資料
func TestQueryAfterInit(t *testing.T) {

	// Arrange
	sacc := new(sa.SimpleAsset)                   // chaincode instance
	stub := shim.NewMockStub("simpleasset", sacc) // shim stub

	key := []byte("k")
	value := []byte("v")

	// Act
	res := stub.MockInit("1", [][]byte{key, value})

	// Init must success
	if int(res.Status) != shim.OK {
		t.FailNow()
	}

	// get key's value
	invokeResult := stub.MockInvoke("1", [][]byte{
		[]byte("get"),
		key,
	})

	// Assert

	// invoke("get") result should equal to value
	assert.Equal(t, invokeResult.Payload, value)
}

// Invoke set a key-value pair should success
func TestInvoke_Set(t *testing.T) {

	// Arrange
	sacc := new(sa.SimpleAsset)                   // chaincode instance
	stub := shim.NewMockStub("simpleasset", sacc) // shim stub

	key := []byte("key123")
	value := []byte("value456")

	// Act
	res := stub.MockInit("1", [][]byte{key, value})

	// Init must success
	if int(res.Status) != shim.OK {
		t.FailNow()
	}

	// Set key value pair
	invokeResult := stub.MockInvoke("1", [][]byte{
		[]byte("set"),
		key,
		value,
	})

	// Assert
	assert.Equal(t, int(invokeResult.Status), shim.OK)
}
