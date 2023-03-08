/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type RequestAttributes struct {
	// access level for setting in module, depending on module it can be string or integer
	AccessLevel string `json:"access_level"`
	// user's id who send request
	FromUser string `json:"from_user"`
	// Submodule where to grant permission
	Link string `json:"link"`
	// Module to grant permission
	Module string `json:"module"`
	// user's id for who request was sent
	ToUser string `json:"to_user"`
	// List of users for whom we give permission
	Users []User `json:"users"`
}
