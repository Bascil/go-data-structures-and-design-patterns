package prototype

import "testing"

func TestClone(t *testing.T){
	shirtCache := GetShirtsCloner()

	if shirtCache == nil {
		t.Fatal("Not implemented yet")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	if item1 == whitePrototype {
		t.Error("Item 1 cannot be equal to white prototype")
	}

	shirt1, ok := item1.(*Shirt)

	if !ok{
		t.Fatal("Type assertion failed for shirt 1")
	}

	shirt1.SKU = "testSKU"

	item2, err := shirtCache.GetClone(White)

    if err != nil {
		t.Fatal(err)
	}

	shirt2, ok := item2.(*Shirt)

	if !ok{
		t.Fatal("Type assertion failed for shirt 2")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("SKUs of shirt 2 and shirt 2 must be different")
	}

	if shirt1 == shirt2 {
		t.Error("Shirt 1 cannot be equal to shirt 2")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())

	t.Logf("LOG:The memory positions of the shirts are different %p != %p\n\n", &shirt1, &shirt2)

}