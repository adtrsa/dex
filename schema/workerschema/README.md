
# Dex API

The Dex REST API

__Version:__ v1

## Models


### AddRemoteIdentityRequest



```
{
    identity: RemoteIdentity
}
```

### AddRemoteIdentityResponse



```
{
    identity: RemoteIdentity
}
```

### DeleteRemoteIdentityRequest



```
{
    identity: RemoteIdentity
}
```

### DeleteRemoteIdentityResponse



```
{
    identity: RemoteIdentity
}
```

### Error



```
{
    error: string,
    error_description: string
}
```

### GetRemoteIdentityResponse



```
{
    identity: RemoteIdentity
}
```

### ListRemoteIdentityResponse



```
{
    identities: [
        RemoteIdentity
    ]
}
```

### RefreshClient

A client with associated public metadata.

```
{
    clientID: string,
    clientName: string,
    clientURI: string,
    logoURI: string
}
```

### RefreshClientList



```
{
    clients: [
        RefreshClient
    ]
}
```

### RemoteIdentity



```
{
    connectorID: string,
    remoteID: string
}
```

### RemoteIdentityDeleteResponse



```
{
    ok: boolean
}
```

### ResendEmailInvitationRequest



```
{
    redirectURL: string
}
```

### ResendEmailInvitationResponse



```
{
    emailSent: boolean,
    resetPasswordLink: string
}
```

### User



```
{
    admin: boolean,
    createdAt: string,
    disabled: boolean,
    displayName: string,
    email: string,
    emailVerified: boolean,
    id: string,
    metadata: string
}
```

### UserCreateRequest



```
{
    redirectURL: string,
    user: User
}
```

### UserCreateResponse



```
{
    emailSent: boolean,
    resetPasswordLink: string,
    user: User
}
```

### UserDeleteResponse



```
{
    ok: boolean
}
```

### UserDisableRequest



```
{
    disable: boolean // If true, disable this user, if false, enable them. No error is signaled if the user state doesn't change.
}
```

### UserDisableResponse



```
{
    ok: boolean
}
```

### UserGetMetadataResponse



```
{
    metadata: string
}
```

### UserResponse



```
{
    user: User
}
```

### UserSetMetadataRequest



```
{
    metadata: string
}
```

### UserSetMetadataResponse



```
{
    ok: boolean
}
```

### UsersResponse



```
{
    nextPageToken: string,
    users: [
        User
    ]
}
```


## Paths


### GET /account/{userid}/refresh

> __Summary__

> List RefreshClient

> __Description__

> List all clients that hold refresh tokens for the specified user.


> __Parameters__

> |Name|Located in|Description|Required|Type|
|:-----|:-----|:-----|:-----|:-----|
| userid | path |  | Yes | string | 


> __Responses__

> |Code|Description|Type|
|:-----|:-----|:-----|
| 200 |  | [RefreshClientList](#refreshclientlist) |
| default | Unexpected error |  |


### DELETE /account/{userid}/refresh/{clientid}

> __Summary__

> Revoke RefreshClient

> __Description__

> Revoke all refresh tokens issues to the client for the specified user.


> __Parameters__

> |Name|Located in|Description|Required|Type|
|:-----|:-----|:-----|:-----|:-----|
| clientid | path |  | Yes | string | 
| userid | path |  | Yes | string | 


> __Responses__

> |Code|Description|Type|
|:-----|:-----|:-----|
| default | Unexpected error |  |


### GET /users

> __Summary__

> List Users

> __Description__

> Retrieve a page of User objects.


> __Parameters__

> |Name|Located in|Description|Required|Type|
|:-----|:-----|:-----|:-----|:-----|
| nextPageToken | query |  | No | string | 
| maxResults | query |  | No | integer | 


> __Responses__

> |Code|Description|Type|
|:-----|:-----|:-----|
| 200 |  | [UsersResponse](#usersresponse) |
| default | Unexpected error |  |


### POST /users

> __Summary__

> Create Users

> __Description__

> Create a new User.


> __Parameters__

> |Name|Located in|Description|Required|Type|
|:-----|:-----|:-----|:-----|:-----|
|  | body |  | Yes | [UserCreateRequest](#usercreaterequest) | 


> __Responses__

> |Code|Description|Type|
|:-----|:-----|:-----|
| 200 |  | [UserCreateResponse](#usercreateresponse) |
| default | Unexpected error |  |


### GET /users/{id}

> __Summary__

> Get Users

> __Description__

> Get a single User object by id.


> __Parameters__

> |Name|Located in|Description|Required|Type|
|:-----|:-----|:-----|:-----|:-----|
| id | path |  | Yes | string | 


> __Responses__

> |Code|Description|Type|
|:-----|:-----|:-----|
| 200 |  | [UserResponse](#userresponse) |
| default | Unexpected error |  |


### POST /users/{id}/disable

> __Summary__

> Disable Users

> __Description__

> Enable or disable a user.


> __Parameters__

> |Name|Located in|Description|Required|Type|
|:-----|:-----|:-----|:-----|:-----|
| id | path |  | Yes | string | 
|  | body |  | Yes | [UserDisableRequest](#userdisablerequest) | 


> __Responses__

> |Code|Description|Type|
|:-----|:-----|:-----|
| 200 |  | [UserDisableResponse](#userdisableresponse) |
| default | Unexpected error |  |


### DELETE /users/{id}/remote-identity

> __Summary__

> DeleteRemoteIdentity Users

> __Description__

> Delete a single RemoteIdentity object by user and supplied remote identity object.


> __Parameters__

> |Name|Located in|Description|Required|Type|
|:-----|:-----|:-----|:-----|:-----|
| id | path |  | Yes | string | 
|  | body |  | Yes | [DeleteRemoteIdentityRequest](#deleteremoteidentityrequest) | 


> __Responses__

> |Code|Description|Type|
|:-----|:-----|:-----|
| 200 |  | [DeleteRemoteIdentityResponse](#deleteremoteidentityresponse) |
| default | Unexpected error |  |


### GET /users/{id}/remote-identity

> __Summary__

> ListRemoteIdentity Users

> __Description__

> Retrieve a page of RemoteIdentity objects.


> __Parameters__

> |Name|Located in|Description|Required|Type|
|:-----|:-----|:-----|:-----|:-----|
| id | path |  | Yes | string | 


> __Responses__

> |Code|Description|Type|
|:-----|:-----|:-----|
| 200 |  | [ListRemoteIdentityResponse](#listremoteidentityresponse) |
| default | Unexpected error |  |


### POST /users/{id}/remote-identity

> __Summary__

> AddRemoteIdentity Users

> __Description__

> Add a remote identity for a user.


> __Parameters__

> |Name|Located in|Description|Required|Type|
|:-----|:-----|:-----|:-----|:-----|
| id | path |  | Yes | string | 
|  | body |  | Yes | [AddRemoteIdentityRequest](#addremoteidentityrequest) | 


> __Responses__

> |Code|Description|Type|
|:-----|:-----|:-----|
| 200 |  | [AddRemoteIdentityResponse](#addremoteidentityresponse) |
| default | Unexpected error |  |


### GET /users/{id}/remote-identity/{connectorid}

> __Summary__

> GetRemoteIdentity Users

> __Description__

> Get a single RemoteIdentity object by user and remote ids.


> __Parameters__

> |Name|Located in|Description|Required|Type|
|:-----|:-----|:-----|:-----|:-----|
| id | path |  | Yes | string | 
| connectorid | path |  | Yes | string | 


> __Responses__

> |Code|Description|Type|
|:-----|:-----|:-----|
| 200 |  | [GetRemoteIdentityResponse](#getremoteidentityresponse) |
| default | Unexpected error |  |


### POST /users/{id}/resend-invitation

> __Summary__

> ResendEmailInvitation Users

> __Description__

> Resend invitation email to an existing user with unverified email.


> __Parameters__

> |Name|Located in|Description|Required|Type|
|:-----|:-----|:-----|:-----|:-----|
| id | path |  | Yes | string | 
|  | body |  | Yes | [ResendEmailInvitationRequest](#resendemailinvitationrequest) | 


> __Responses__

> |Code|Description|Type|
|:-----|:-----|:-----|
| 200 |  | [ResendEmailInvitationResponse](#resendemailinvitationresponse) |
| default | Unexpected error |  |


