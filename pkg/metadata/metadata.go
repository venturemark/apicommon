// Package metadata provides label constants used for resource identification
// and information sharing. Below is an incomplete list of resources we may take
// care of.
//
//     message
//     timeline
//     update
//     user
//
// The above resources have different properties. It might be necessary to
// expose different pieces of information throughout the lifecycle of the
// resources when interacting with them.
//
//     id        unix timestamp normalized to UTC timezone
//     status    state change information e.g. "updated"
//
package metadata
