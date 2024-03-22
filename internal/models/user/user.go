package user

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrSerializationError 	= errors.New("serialization error")
	ErrDeserializationError = errors.New("deserialization error")
)

type User struct {

	// Id of the chat.
	ID 		 	string 		

	// Name of the user.
	UserName 	string 	

	// Active status of the action.	
	IsCreate 	bool

	IsAdd 	 	bool   		

	IsGet 	 	bool   		

	IsDelete 	bool

	// Emoji of the sticker.
	Emoji 	 	string 		

	// Name of the sticker set.
	SetName  	string 		

	// Bytes of the User structure.
	ByteUser 	[]byte

	// Iterator for the action.
	IterAction 	int
}
// Serialize encodes the user into a byte slice.
func (u *User) Serialize() (err error){  
	u.ByteUser = []byte(u.Prepare()) 
	if u.ByteUser == nil {
		return ErrSerializationError
	}
	return 
}
// Deserialize decodes the user from a byte slice.
func (u *User) Deserialize() (err error) {  		
	str := strings.Split(string(u.ByteUser), " ")
	u.ID = str[0]
	u.UserName = str[1]
	u.IsCreate, err = strconv.ParseBool(str[2])
	if err!= nil {
        return ErrDeserializationError
    }
	u.IsAdd, err = strconv.ParseBool(str[3])
	if err!= nil {
        return ErrDeserializationError
    }
	u.IsGet, err = strconv.ParseBool(str[4])
	if err!= nil {
        return ErrDeserializationError
    }
	u.IsDelete, err = strconv.ParseBool(str[5])
	if err!= nil {
        return ErrDeserializationError
    }
	u.Emoji = str[6]
	u.SetName = str[7]
	u.IterAction, err = strconv.Atoi(str[8])
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Prepare() string{
	return fmt.Sprintf("%s %s %t %t %t %t %s %s %d",
		    u.ID, u.UserName, u.IsCreate, u.IsAdd, u.IsGet, u.IsDelete, u.Emoji, u.SetName, u.IterAction)
}

func (c *User) Recipient() string{
	return c.ID
}