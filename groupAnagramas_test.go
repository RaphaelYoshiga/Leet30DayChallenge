package main

import "testing"
import "strings"

func TestAnagramGroupping(t *testing.T) {
    tables := []struct {
        x []string;
        exp [][]string
    }{
        {[]string {}, [][]string {}},
        {[]string { "test"}, [][]string { {"test"}} },
        {[]string { "test", "lel"}, [][]string { {"test"}, {"lel"}} },
        {[]string { "nop", "nap"}, [][]string { {"nop"}, {"nap"}} },
        {[]string { "nop", "nop", "nap"}, [][]string { {"nop", "nop"}, {"nap"}} },
        {[]string { "eat", "tea", "tan", "ate", "nat", "bat"}, [][]string { {"eat", "tea", "ate"}, {"tan", "nat"}, { "bat"}} },
    }
    
    for _, table := range tables {

        result := groupAnagrams(table.x);

        for i, line := range table.exp {
            for y, word := range line{
                actual:= result[i][y];
                if word != actual {
                    t.Errorf("[%s] Anagram not found %s, was %s", strings.Join(table.x, ","), word, actual)
                    break;
                 }
            }
        }
    }
}


func TestHash(t *testing.T) {
    tables := []struct {
        a string;
        b string
    }{
        {"test", "estt" },
        {"lapa", "pala" },
    }
    
    for _, table := range tables {

        if hash(table.a) != hash(table.b) {
            t.Errorf("%s should be same hash of %s", table.a, table.b)
        }
    }
}

func TestHashesNotMatch(t *testing.T) {
    tables := []struct {
        a string;
        b string
    }{
        {"ttt", "aaa" },
        {"ddd", "fff" },
    }
    
    for _, table := range tables {

        if hash(table.a) == hash(table.b) {
            t.Errorf("%s should be same hash of %s", table.a, table.b)
        }
    }
}