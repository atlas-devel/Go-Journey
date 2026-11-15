package main

type Notifier interface{
	send(message string)
}