package problem_test

import (
	"testing"
	"time"

	"github.com/AlexsJones/devops-go-problems/bit.ly/problem"
)

const testURL = "www.beamery.com"

func TestBasicStorer(t *testing.T) {
	storer := problem.NewStorer(1 * time.Second)
	if storer == nil {
		t.Fatal("Unimplemented store")
	}
	l, err := storer.AddLink(testURL)
	if err != nil {
		t.Fatal("Should not return an error at this point")
	}
	if l.GetURL() != testURL {
		t.Fatal("Does not store the correct url")
	}
	if l.Expired() {
		t.Fatal("Link should not have expired")
	}
	<-time.After(1 * time.Second)
	if !l.Expired() {
		t.Fatal("Link should of expired by now")
	}
}

func TestStorerMaintance(t *testing.T) {
	storer := problem.NewStorer(1 * time.Second)
	if storer == nil {
		t.Fatal("Unimplemented store")
	}
	l, err := storer.AddLink(testURL)
	if err != nil {
		t.Fatal("Should not return an error at this point")
	}
	if l.GetURL() != testURL {
		t.Fatal("Does not store the correct url")
	}
	if l.Expired() {
		t.Fatal("Link should not have expired")
	}
	if _, err = storer.AddLink(l.GetURL()); err == nil {
		t.Fatal("Link had already been added and should report an error")
	}
	link, err := storer.GetLink(l.GetShortLink())
	if err != nil {
		t.Fatal("Link should exist")
	}
	if link.GetShortLink() != l.GetShortLink() {
		t.Fatal("The links should be equal")
	}
	if _, err = storer.GetLink(testURL); err == nil {
		t.Fatal("Used url to look up link instead of shotern link")
	}
	if l != link {
		t.Fatal("Should be the same object")
	}
	<-time.After(1 * time.Second)
	if !l.Expired() {
		t.Fatal("Link should have expired")
	}
	if _, err = storer.AddLink(testURL); err == nil {
		t.Fatal("Storer should still have reference to expired objects")
	}
}
