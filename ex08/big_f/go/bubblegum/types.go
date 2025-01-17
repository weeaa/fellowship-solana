// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package bubblegum

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Creator struct {
	Address  ag_solanago.PublicKey
	Verified bool
	Share    uint8
}

func (obj Creator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `Verified` param:
	err = encoder.Encode(obj.Verified)
	if err != nil {
		return err
	}
	// Serialize `Share` param:
	err = encoder.Encode(obj.Share)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Creator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `Verified`:
	err = decoder.Decode(&obj.Verified)
	if err != nil {
		return err
	}
	// Deserialize `Share`:
	err = decoder.Decode(&obj.Share)
	if err != nil {
		return err
	}
	return nil
}

type Uses struct {
	UseMethod UseMethod
	Remaining uint64
	Total     uint64
}

func (obj Uses) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UseMethod` param:
	err = encoder.Encode(obj.UseMethod)
	if err != nil {
		return err
	}
	// Serialize `Remaining` param:
	err = encoder.Encode(obj.Remaining)
	if err != nil {
		return err
	}
	// Serialize `Total` param:
	err = encoder.Encode(obj.Total)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Uses) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UseMethod`:
	err = decoder.Decode(&obj.UseMethod)
	if err != nil {
		return err
	}
	// Deserialize `Remaining`:
	err = decoder.Decode(&obj.Remaining)
	if err != nil {
		return err
	}
	// Deserialize `Total`:
	err = decoder.Decode(&obj.Total)
	if err != nil {
		return err
	}
	return nil
}

type Collection struct {
	Verified bool
	Key      ag_solanago.PublicKey
}

func (obj Collection) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Verified` param:
	err = encoder.Encode(obj.Verified)
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Collection) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Verified`:
	err = decoder.Decode(&obj.Verified)
	if err != nil {
		return err
	}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	return nil
}

type MetadataArgs struct {
	Name                 string
	Symbol               string
	Uri                  string
	SellerFeeBasisPoints uint16
	PrimarySaleHappened  bool
	IsMutable            bool
	EditionNonce         *uint8         `bin:"optional"`
	TokenStandard        *TokenStandard `bin:"optional"`
	Collection           *Collection    `bin:"optional"`
	Uses                 *Uses          `bin:"optional"`
	TokenProgramVersion  TokenProgramVersion
	Creators             []Creator
}

func (obj MetadataArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `Uri` param:
	err = encoder.Encode(obj.Uri)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `PrimarySaleHappened` param:
	err = encoder.Encode(obj.PrimarySaleHappened)
	if err != nil {
		return err
	}
	// Serialize `IsMutable` param:
	err = encoder.Encode(obj.IsMutable)
	if err != nil {
		return err
	}
	// Serialize `EditionNonce` param (optional):
	{
		if obj.EditionNonce == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.EditionNonce)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `TokenStandard` param (optional):
	{
		if obj.TokenStandard == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TokenStandard)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Collection` param (optional):
	{
		if obj.Collection == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Uses` param (optional):
	{
		if obj.Uses == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Uses)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `TokenProgramVersion` param:
	err = encoder.Encode(obj.TokenProgramVersion)
	if err != nil {
		return err
	}
	// Serialize `Creators` param:
	err = encoder.Encode(obj.Creators)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MetadataArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `Uri`:
	err = decoder.Decode(&obj.Uri)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `PrimarySaleHappened`:
	err = decoder.Decode(&obj.PrimarySaleHappened)
	if err != nil {
		return err
	}
	// Deserialize `IsMutable`:
	err = decoder.Decode(&obj.IsMutable)
	if err != nil {
		return err
	}
	// Deserialize `EditionNonce` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.EditionNonce)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `TokenStandard` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TokenStandard)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Collection` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Uses` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Uses)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `TokenProgramVersion`:
	err = decoder.Decode(&obj.TokenProgramVersion)
	if err != nil {
		return err
	}
	// Deserialize `Creators`:
	err = decoder.Decode(&obj.Creators)
	if err != nil {
		return err
	}
	return nil
}

type UpdateArgs struct {
	Name                 *string    `bin:"optional"`
	Symbol               *string    `bin:"optional"`
	Uri                  *string    `bin:"optional"`
	Creators             *[]Creator `bin:"optional"`
	SellerFeeBasisPoints *uint16    `bin:"optional"`
	PrimarySaleHappened  *bool      `bin:"optional"`
	IsMutable            *bool      `bin:"optional"`
}

func (obj UpdateArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param (optional):
	{
		if obj.Name == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Name)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Symbol` param (optional):
	{
		if obj.Symbol == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Symbol)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Uri` param (optional):
	{
		if obj.Uri == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Uri)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Creators` param (optional):
	{
		if obj.Creators == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `SellerFeeBasisPoints` param (optional):
	{
		if obj.SellerFeeBasisPoints == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SellerFeeBasisPoints)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `PrimarySaleHappened` param (optional):
	{
		if obj.PrimarySaleHappened == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.PrimarySaleHappened)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `IsMutable` param (optional):
	{
		if obj.IsMutable == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.IsMutable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *UpdateArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Name)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Symbol` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Symbol)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Uri` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Uri)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Creators` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `SellerFeeBasisPoints` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SellerFeeBasisPoints)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `PrimarySaleHappened` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.PrimarySaleHappened)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `IsMutable` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.IsMutable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Version ag_binary.BorshEnum

const (
	VersionV1 Version = iota
)

func (value Version) String() string {
	switch value {
	case VersionV1:
		return "V1"
	default:
		return ""
	}
}

type LeafSchema interface {
	isLeafSchema()
}

type leafSchemaContainer struct {
	Enum ag_binary.BorshEnum `borsh_enum:"true"`
	V1   LeafSchemaV1
}

type LeafSchemaV1 struct {
	Id          ag_solanago.PublicKey
	Owner       ag_solanago.PublicKey
	Delegate    ag_solanago.PublicKey
	Nonce       uint64
	DataHash    [32]uint8
	CreatorHash [32]uint8
}

func (obj LeafSchemaV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Id` param:
	err = encoder.Encode(obj.Id)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `Delegate` param:
	err = encoder.Encode(obj.Delegate)
	if err != nil {
		return err
	}
	// Serialize `Nonce` param:
	err = encoder.Encode(obj.Nonce)
	if err != nil {
		return err
	}
	// Serialize `DataHash` param:
	err = encoder.Encode(obj.DataHash)
	if err != nil {
		return err
	}
	// Serialize `CreatorHash` param:
	err = encoder.Encode(obj.CreatorHash)
	if err != nil {
		return err
	}
	return nil
}

func (obj *LeafSchemaV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Id`:
	err = decoder.Decode(&obj.Id)
	if err != nil {
		return err
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `Delegate`:
	err = decoder.Decode(&obj.Delegate)
	if err != nil {
		return err
	}
	// Deserialize `Nonce`:
	err = decoder.Decode(&obj.Nonce)
	if err != nil {
		return err
	}
	// Deserialize `DataHash`:
	err = decoder.Decode(&obj.DataHash)
	if err != nil {
		return err
	}
	// Deserialize `CreatorHash`:
	err = decoder.Decode(&obj.CreatorHash)
	if err != nil {
		return err
	}
	return nil
}

func (_ *LeafSchemaV1) isLeafSchema() {}

type TokenProgramVersion ag_binary.BorshEnum

const (
	TokenProgramVersionOriginal TokenProgramVersion = iota
	TokenProgramVersionToken2022
)

func (value TokenProgramVersion) String() string {
	switch value {
	case TokenProgramVersionOriginal:
		return "Original"
	case TokenProgramVersionToken2022:
		return "Token2022"
	default:
		return ""
	}
}

type TokenStandard ag_binary.BorshEnum

const (
	TokenStandardNonFungible TokenStandard = iota
	TokenStandardFungibleAsset
	TokenStandardFungible
	TokenStandardNonFungibleEdition
)

func (value TokenStandard) String() string {
	switch value {
	case TokenStandardNonFungible:
		return "NonFungible"
	case TokenStandardFungibleAsset:
		return "FungibleAsset"
	case TokenStandardFungible:
		return "Fungible"
	case TokenStandardNonFungibleEdition:
		return "NonFungibleEdition"
	default:
		return ""
	}
}

type UseMethod ag_binary.BorshEnum

const (
	UseMethodBurn UseMethod = iota
	UseMethodMultiple
	UseMethodSingle
)

func (value UseMethod) String() string {
	switch value {
	case UseMethodBurn:
		return "Burn"
	case UseMethodMultiple:
		return "Multiple"
	case UseMethodSingle:
		return "Single"
	default:
		return ""
	}
}

type BubblegumEventType ag_binary.BorshEnum

const (
	BubblegumEventTypeUninitialized BubblegumEventType = iota
	BubblegumEventTypeLeafSchemaEvent
)

func (value BubblegumEventType) String() string {
	switch value {
	case BubblegumEventTypeUninitialized:
		return "Uninitialized"
	case BubblegumEventTypeLeafSchemaEvent:
		return "LeafSchemaEvent"
	default:
		return ""
	}
}

type DecompressibleState ag_binary.BorshEnum

const (
	DecompressibleStateEnabled DecompressibleState = iota
	DecompressibleStateDisabled
)

func (value DecompressibleState) String() string {
	switch value {
	case DecompressibleStateEnabled:
		return "Enabled"
	case DecompressibleStateDisabled:
		return "Disabled"
	default:
		return ""
	}
}
