# Duplicated Code

Two code fragments look almost identical.


## What It Looks Like

```
type Request struct {
    ID int64
    Description string
}

func (req *Request) create(userRole string) error {
    if userRole == "ADMIN" {
        Err("This role is not allowed to do this action")
    }

    // logic to create request
}

func (req *Request) delete(userRole string) error {
    if userRole == "ADMIN" {
        return Err("This role is not allowed to do this action")
    }

    // logic to delete request
}
```

## Why It Hurts

It can lead to maintenance challenges, increased risk of errors, and reduced code maintainability.

## How To Fix It

- If the same code is found in two or more methods in the same class: use [Extract Method](.././../2.%20refactorings/extract-method.md).

## Refactor

```
type Request struct {
    ID          int64
    Description string
}

func (req *Request) checkPermission(userRole string) error {
    if userRole == "ADMIN" {
        return Err("This role is not allowed to do this action")
    }
    return nil
}

func (req *Request) create(userRole string) error {
    if err := req.checkPermission(userRole); err != nil {
        return err
    }

    // Logic to create request
    return nil
}

func (req *Request) delete(userRole string) error {
    if err := req.checkPermission(userRole); err != nil {
        return err
    }

    // Logic to delete request
    return nil
}

```

## Payoff

- Merging duplicate code simplifies the structure of your code and makes it shorter.
- Simplification + shortness = code that’s easier to simplify and cheaper to support.

