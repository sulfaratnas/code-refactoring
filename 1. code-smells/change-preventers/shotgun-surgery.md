# Shotgun Surgery

The "Shotgun Surgery" code smell occurs when a single change to a feature or functionality requires modifications in multiple places across the codebase. 

## What It Looks Like

```
package main

import "fmt"

type userRole int

const ADMIN_ROLE userRole = 1

type User struct {
 role userRole
}

func (u *User) Create() {
 if u.role == ADMIN_ROLE {
  fmt.Println("User is allowed to create request")
 }
}

func (u *User) Update() {
 if u.role == ADMIN_ROLE {
  fmt.Println("User is allowed to update request")
 }
}

func (u *User) Delete() {
 if u.role == ADMIN_ROLE {
  fmt.Println("User is allowed to delete request")
 }
}
```
In this example, we have three methods. Each method checks the user's role. Whenever we need to modify the logic for checking the user's role or if we want to rename the const ADMIN_ROLE it will impact 3 methods and we find ourselves making changes in multiple places, which can result in the "Shotgun Surgery" code smell.

## Why It Hurts

- The more places you have to make changes, the greater the chance of introducing errors.
- Shotgun surgery often leads to changes being scattered across multiple files or classes.
- Frequent shotgun surgery can significantly increase the time it takes to implement new features or fix bugs. 
- Decreased Code Readability.

## How To Fix It

Use [Move Method](.././../2.%20refactorings/move-method.md) to move existing class behaviors into a single class. If there’s no class appropriate for this, create a new one.

If moving code to the same class leaves the original classes almost empty, try to get rid of these now-redundant classes via [Inline Class](.././../2.%20refactorings/inline-class.md).

## Refactor

we need to create a separate method and that method will used by the three methods. Let’s see the improvement code here:
```
package main

import "fmt"

type userRole int

const ADMIN_ROLE userRole = 1

type User struct {
 role userRole
}

func (u *User) Create() {
 if u.isUserAdmin() {
  fmt.Println("User is allowed to create request")
 }
}

func (u *User) Update() {
 if u.isUserAdmin() {
  fmt.Println("User is allowed to update request")
 }
}

func (u *User) Delete() {
 if u.isUserAdmin() {
  fmt.Println("User is allowed to delete request")
 }
}

func (u *User) isUserAdmin() bool {
 return u.role == ADMIN_ROLE
}
```



## Payoff

- Better organization.
- Less code duplication.
- Easier maintenance.
