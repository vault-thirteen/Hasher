package ot

type ObjectTypeName string

const (
	Name_File      = ObjectTypeName("FILE")
	Name_Folder    = ObjectTypeName("FOLDER")
	Name_Directory = ObjectTypeName("DIRECTORY")
)
