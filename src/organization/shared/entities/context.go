// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
)

type clientCtxKey struct{}

// FromContext returns a Client stored inside a context, or nil if there isn't one.
func FromContext(ctx context.Context) *Client {
	c, _ := ctx.Value(clientCtxKey{}).(*Client)
	return c
}

// NewContext returns a new context with the given Client attached.
func NewContext(parent context.Context, c *Client) context.Context {
	return context.WithValue(parent, clientCtxKey{}, c)
}

type txCtxKey struct{}

// TxFromContext returns a Tx stored inside a context, or nil if there isn't one.
func TxFromContext(ctx context.Context) *Tx {
	tx, _ := ctx.Value(txCtxKey{}).(*Tx)
	return tx
}

// NewTxContext returns a new context with the given Tx attached.
func NewTxContext(parent context.Context, tx *Tx) context.Context {
	return context.WithValue(parent, txCtxKey{}, tx)
}
