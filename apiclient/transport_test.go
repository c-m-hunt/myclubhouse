package apiclient_test

import (
	"testing"

	"github.com/c-m-hunt/myclubhouse/apiclient"
)

func TestItGetsNextPageFromFirstPage(t *testing.T) {
	q := apiclient.RequestQuery{
		PageSize: 10,
	}
	rp := apiclient.ResponsePagination{
		ItemCount:    33,
		PageCount:    4,
		PageSize:     10,
		SelectedPage: 0,
	}

	q.NextPage(rp)

	if q.SelectedPage != 1 {
		t.Fatal("Next page does not increment from 0")
	}

}

func TestItGetsNextPageFromLaterPages(t *testing.T) {

	q := apiclient.RequestQuery{
		PageSize:     10,
		SelectedPage: 2,
	}
	rp := apiclient.ResponsePagination{
		ItemCount:    33,
		PageCount:    4,
		PageSize:     10,
		SelectedPage: 2,
	}
	q.NextPage(rp)

	if q.SelectedPage != 3 {
		t.Fatal("Next page does not increment from 2")
	}
}

func TestItReturnsAnErrorIfTheresNoMorePages(t *testing.T) {

	q := apiclient.RequestQuery{
		PageSize:     10,
		SelectedPage: 3,
	}
	rp := apiclient.ResponsePagination{
		ItemCount:    33,
		PageCount:    4,
		PageSize:     10,
		SelectedPage: 3,
	}
	err := q.NextPage(rp)

	if err == nil {
		t.Fatal("Next page does not know when it's the end")
	}
}
