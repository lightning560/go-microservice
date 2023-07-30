// Code generated by ent, DO NOT EDIT.

package ent

import (
	"redbook/app/domain/user/internal/data/ent/profile"
	"redbook/app/domain/user/internal/data/ent/schema"
	"redbook/app/domain/user/internal/data/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescSex is the schema descriptor for sex field.
	profileDescSex := profileFields[3].Descriptor()
	// profile.DefaultSex holds the default value on creation for the sex field.
	profile.DefaultSex = profileDescSex.Default.(int32)
	// profileDescAvatar is the schema descriptor for avatar field.
	profileDescAvatar := profileFields[4].Descriptor()
	// profile.DefaultAvatar holds the default value on creation for the avatar field.
	profile.DefaultAvatar = profileDescAvatar.Default.(string)
	// profileDescLevel is the schema descriptor for level field.
	profileDescLevel := profileFields[7].Descriptor()
	// profile.DefaultLevel holds the default value on creation for the level field.
	profile.DefaultLevel = profileDescLevel.Default.(int32)
	// profileDescVerifyType is the schema descriptor for verify_type field.
	profileDescVerifyType := profileFields[8].Descriptor()
	// profile.DefaultVerifyType holds the default value on creation for the verify_type field.
	profile.DefaultVerifyType = profileDescVerifyType.Default.(int8)
	// profileDescAttr is the schema descriptor for attr field.
	profileDescAttr := profileFields[9].Descriptor()
	// profile.DefaultAttr holds the default value on creation for the attr field.
	profile.DefaultAttr = profileDescAttr.Default.(int32)
	// profileDescState is the schema descriptor for state field.
	profileDescState := profileFields[10].Descriptor()
	// profile.DefaultState holds the default value on creation for the state field.
	profile.DefaultState = profileDescState.Default.(int32)
	// profileDescCreatedAt is the schema descriptor for created_at field.
	profileDescCreatedAt := profileFields[12].Descriptor()
	// profile.DefaultCreatedAt holds the default value on creation for the created_at field.
	profile.DefaultCreatedAt = profileDescCreatedAt.Default.(func() time.Time)
	// profileDescUpdatedAt is the schema descriptor for updated_at field.
	profileDescUpdatedAt := profileFields[13].Descriptor()
	// profile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	profile.DefaultUpdatedAt = profileDescUpdatedAt.Default.(func() time.Time)
	// profile.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	profile.UpdateDefaultUpdatedAt = profileDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUID is the schema descriptor for uid field.
	userDescUID := userFields[1].Descriptor()
	// user.UIDValidator is a validator for the "uid" field. It is called by the builders before save.
	user.UIDValidator = userDescUID.Validators[0].(func(int64) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[7].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[8].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(int64) error)
}
