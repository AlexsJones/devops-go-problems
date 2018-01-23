# Bit.ly Service Problem

## Problem Definition

The domain [bit ly](https://bitly.com) offers a service that allows their users
to give them a URL to shorten and be accessible given a unique link that they give back.

We want to reimplement this service in GoLang (because it fair superior than other system languages).
The constraints we want to enforce are the following:

- [ ] Given a URL
  - When you apply our _"shorten"_ method to it
  - Then it should always return the same shorten link.
- [ ] Given a new shorten link
  - And the client hasn't paid
  -  Then it should expire after X units of time
- [ ] Given a shorten link
    - And the client has paid
    - Then the link should never expire

## Extension Problem
- [ ] Given that we will need to upgrade
  - And keep all the current data stored
  - Then we should be able to upgrade seamlessly

## Where to start
In order to start working on this problem, we have written some [boiler plate code](problem/) which should be enough to get you started.

We have also added the ability to test your implementation by _actually_ using your bitly service.

The service address is: `127.0.0.1:8000`
The paths it supports are:
- `/get/{link}`
  - Method: `GET`
  - link : the shorten link produced by your storer.
- `/add/{link}`
  - Method: `POST`
  - link : The link you want to shorten.

## Testing your code
We have written some [tests](test/) to ensure that you have implemented a solution to our problem correctly.

```sh
# To run all of our test cases
go test -v ./...
```
