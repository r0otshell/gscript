// Package engine is the package that implements the gscript Virtual Machine (VM).
//
// VM Functions implemented:
//
// Library core
//
// Functions in core:
//  Asset(assetName) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.Asset
//  B64Decode(data) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.B64Decode
//  B64Encode(data) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.B64Encode
//  DeobfuscateString(str) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.DeobfuscateString
//  Halt() - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.Halt
//  MD5(data) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.MD5
//  ObfuscateString(str) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.ObfuscateString
//  RandomInt(min, max) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.RandomInt
//  RandomMixedCaseString(strlen) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.RandomMixedCaseString
//  RandomString(strlen) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.RandomString
//  SHA1(data) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.SHA1
//  StripSpaces(str) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.StripSpaces
//  Timestamp() - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.Timestamp
//  XorBytes(aByteArray, bByteArray) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.XorBytes
//
// Library exec
//
// Functions in exec:
//  ExecuteCommand(baseCmd, cmdArgs) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.ExecuteCommand
//  ForkExecuteCommand(baseCmd, cmdArgs) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.ForkExecuteCommand
//
// Library file
//
// Functions in file:
//  AppendFileBytes(path, fileData) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.AppendFileBytes
//  AppendFileString(path, addString) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.AppendFileString
//  CopyFile(srcPath, dstPath, perms) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.CopyFile
//  CreateDir(path) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.CreateDir
//  DeleteFile(path) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.DeleteFile
//  FileExists(path) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.FileExists
//  ReadFile(path) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.ReadFile
//  ReplaceFileString(file, match, replacement) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.ReplaceFileString
//  WriteFile(path, fileData, perms) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.WriteFile
//  WriteTempFile(name, fileData) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.WriteTempFile
//
// Library injection
//
// Functions in injection:
//  InjectIntoProc(shellcode, proccessID) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.InjectIntoProc
//  InjectIntoSelf(shellcode) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.InjectIntoSelf
//
// Library net
//
// Functions in net:
//  DNSQuestion(target, request) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.DNSQuestion
//  GetLocalIPs() - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.GetLocalIPs
//  GetMACAddress() - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.GetMACAddress
//  HTTPGetFile(url) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.HTTPGetFile
//  IsTCPPortInUse(port) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.IsTCPPortInUse
//  IsUDPPortInUse(port) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.IsUDPPortInUse
//  PostJSON(url, json) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.PostJSON
//  SSHCmd(hostAndPort, cmd, username, password, key) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.SSHCmd
//  ServePathOverHTTPS(port, path, timeout) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.ServePathOverHTTPS
//
// Library os
//
// Functions in os:
//  Chmod(path, perms) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.Chmod
//  EnvVars() - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.EnvVars
//  FindProcByName(procName) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.FindProcByName
//  GetEnvVar(vars) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.GetEnvVar
//  GetProcName(pid) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.GetProcName
//  InstallSystemService(path, name, displayName, description) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.InstallSystemService
//  ModTime(path) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.ModTime
//  ModifyTimestamp(path, accessTime, modifyTime) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.ModifyTimestamp
//  RemoveServiceByName(name) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.RemoveServiceByName
//  RunningProcs() - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.RunningProcs
//  SelfPath() - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.SelfPath
//  Signal(signal, pid) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.Signal
//  StartServiceByName(name) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.StartServiceByName
//  StopServiceByName(name) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.StopServiceByName
//
// Library polymorph
//
// Functions in polymorph:
//  RetrievePEPolymorphicData(peFile) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.RetrievePEPolymorphicData
//  WritePEPolymorphicData(peFile, data) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.WritePEPolymorphicData
//
// Library registry
//
// Functions in registry:
//  AddRegKeyBinary(registryString, path, name, value) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.AddRegKeyBinary
//  AddRegKeyDWORD(registryString, path, name, value) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.AddRegKeyDWORD
//  AddRegKeyExpandedString(registryString, path, name, value) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.AddRegKeyExpandedString
//  AddRegKeyQWORD(registryString, path, name, value) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.AddRegKeyQWORD
//  AddRegKeyString(registryString, path, name, value) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.AddRegKeyString
//  AddRegKeyStrings(registryString, path, name, value) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.AddRegKeyStrings
//  DelRegKey(registryString, path) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.DelRegKey
//  DelRegKeyValue(registryString, path, value) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.DelRegKeyValue
//  QueryRegKey(registryString, path) - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.QueryRegKey
//
// Library stealth
//
// Functions in stealth:
//  CheckIfCPUCountIsHigherThanOne() - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.CheckIfCPUCountIsHigherThanOne
//  CheckIfRAMAmountIsBelow1GB() - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.CheckIfRAMAmountIsBelow1GB
//  CheckSandboxUsernames() - https://godoc.org/github.com/gen0cide/gscript/engine/#Engine.CheckSandboxUsernames
//
package engine

import (
	"path/filepath"

	"github.com/robertkrimen/otto"
)

func (e *Engine) CreateVM() {
	e.VM = otto.New()
	e.injectVars()
	e.VM.Set("AddRegKeyBinary", e.vmAddRegKeyBinary)
	e.VM.Set("AddRegKeyDWORD", e.vmAddRegKeyDWORD)
	e.VM.Set("AddRegKeyExpandedString", e.vmAddRegKeyExpandedString)
	e.VM.Set("AddRegKeyQWORD", e.vmAddRegKeyQWORD)
	e.VM.Set("AddRegKeyString", e.vmAddRegKeyString)
	e.VM.Set("AddRegKeyStrings", e.vmAddRegKeyStrings)
	e.VM.Set("AppendFileBytes", e.vmAppendFileBytes)
	e.VM.Set("AppendFileString", e.vmAppendFileString)
	e.VM.Set("Asset", e.vmAsset)
	e.VM.Set("B64Decode", e.vmB64Decode)
	e.VM.Set("B64Encode", e.vmB64Encode)
	e.VM.Set("CheckIfCPUCountIsHigherThanOne", e.vmCheckIfCPUCountIsHigherThanOne)
	e.VM.Set("CheckIfRAMAmountIsBelow1GB", e.vmCheckIfRAMAmountIsBelow1GB)
	e.VM.Set("CheckIfWineGetUnixFileNameExists", e.vmCheckIfWineGetUnixFileNameExists)
	e.VM.Set("CheckSandboxUsernames", e.vmCheckSandboxUsernames)
	e.VM.Set("Chmod", e.vmChmod)
	e.VM.Set("CopyFile", e.vmCopyFile)
	e.VM.Set("CreateDir", e.vmCreateDir)
	e.VM.Set("DNSQuestion", e.vmDNSQuestion)
	e.VM.Set("DelRegKey", e.vmDelRegKey)
	e.VM.Set("DelRegKeyValue", e.vmDelRegKeyValue)
	e.VM.Set("DeleteFile", e.vmDeleteFile)
	e.VM.Set("DeobfuscateString", e.vmDeobfuscateString)
	e.VM.Set("EnvVars", e.vmEnvVars)
	e.VM.Set("ExecuteCommand", e.vmExecuteCommand)
	e.VM.Set("FileExists", e.vmFileExists)
	e.VM.Set("FindProcByName", e.vmFindProcByName)
	e.VM.Set("ForkExecuteCommand", e.vmForkExecuteCommand)
	e.VM.Set("GetEnvVar", e.vmGetEnvVar)
	e.VM.Set("GetLocalIPs", e.vmGetLocalIPs)
	e.VM.Set("GetMACAddress", e.vmGetMACAddress)
	e.VM.Set("GetProcName", e.vmGetProcName)
	e.VM.Set("HTTPGetFile", e.vmHTTPGetFile)
	e.VM.Set("Halt", e.vmHalt)
	e.VM.Set("InjectIntoProc", e.vmInjectIntoProc)
	e.VM.Set("InjectIntoSelf", e.vmInjectIntoSelf)
	e.VM.Set("InstallSystemService", e.vmInstallSystemService)
	e.VM.Set("IsDebuggerPresent", e.vmIsDebuggerPresent)
	e.VM.Set("IsTCPPortInUse", e.vmIsTCPPortInUse)
	e.VM.Set("IsUDPPortInUse", e.vmIsUDPPortInUse)
	e.VM.Set("MD5", e.vmMD5)
	e.VM.Set("MakeDebuggable", e.vmMakeDebuggable)
	e.VM.Set("MakeUnDebuggable", e.vmMakeUnDebuggable)
	e.VM.Set("ModTime", e.vmModTime)
	e.VM.Set("ModifyTimestamp", e.vmModifyTimestamp)
	e.VM.Set("ObfuscateString", e.vmObfuscateString)
	e.VM.Set("PostJSON", e.vmPostJSON)
	e.VM.Set("QueryRegKey", e.vmQueryRegKey)
	e.VM.Set("RandomInt", e.vmRandomInt)
	e.VM.Set("RandomMixedCaseString", e.vmRandomMixedCaseString)
	e.VM.Set("RandomString", e.vmRandomString)
	e.VM.Set("ReadFile", e.vmReadFile)
	e.VM.Set("RemoveServiceByName", e.vmRemoveServiceByName)
	e.VM.Set("ReplaceFileString", e.vmReplaceFileString)
	e.VM.Set("RetrievePEPolymorphicData", e.vmRetrievePEPolymorphicData)
	e.VM.Set("RunningProcs", e.vmRunningProcs)
	e.VM.Set("SHA1", e.vmSHA1)
	e.VM.Set("SSHCmd", e.vmSSHCmd)
	e.VM.Set("SelfPath", e.vmSelfPath)
	e.VM.Set("ServePathOverHTTPS", e.vmServePathOverHTTPS)
	e.VM.Set("Signal", e.vmSignal)
	e.VM.Set("StartServiceByName", e.vmStartServiceByName)
	e.VM.Set("StopServiceByName", e.vmStopServiceByName)
	e.VM.Set("StripSpaces", e.vmStripSpaces)
	e.VM.Set("Timestamp", e.vmTimestamp)
	e.VM.Set("WriteFile", e.vmWriteFile)
	e.VM.Set("WritePEPolymorphicData", e.vmWritePEPolymorphicData)
	e.VM.Set("WriteTempFile", e.vmWriteTempFile)
	e.VM.Set("XorBytes", e.vmXorBytes)
	_, err := e.VM.Run(vmPreload)
	if err != nil {
		e.Logger.WithField("trace", "true").Fatalf("Syntax error in preload: %s", err.Error())
	}
	e.initializeLogger()
}

func (e *Engine) vmAddRegKeyBinary(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 4 {
		e.Logger.WithField("function", "AddRegKeyBinary").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 4 {
		e.Logger.WithField("function", "AddRegKeyBinary").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var registryString string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyBinary").WithField("trace", "true").Errorf("Could not export field: %s", "registryString")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		registryString = rawArg0.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyBinary").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var path string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyBinary").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		path = filepath.Clean(rawArg1.(string))
	default:
		e.Logger.WithField("function", "AddRegKeyBinary").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var name string
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyBinary").WithField("trace", "true").Errorf("Could not export field: %s", "name")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case string:
		name = rawArg2.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyBinary").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	value := e.ValueToByteSlice(call.Argument(3))
	runtimeError := e.AddRegKeyBinary(registryString, path, name, value)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "AddRegKeyBinary").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "AddRegKeyBinary").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmAddRegKeyDWORD(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 4 {
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 4 {
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var registryString string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Errorf("Could not export field: %s", "registryString")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		registryString = rawArg0.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var path string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		path = filepath.Clean(rawArg1.(string))
	default:
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var name string
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Errorf("Could not export field: %s", "name")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case string:
		name = rawArg2.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var value uint32
	rawArg3, err := call.Argument(3).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Errorf("Could not export field: %s", "value")
		return otto.FalseValue()
	}
	switch v := rawArg3.(type) {
	case uint32:
		value = rawArg3.(uint32)
	default:
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "uint32", v)
		return otto.FalseValue()
	}
	runtimeError := e.AddRegKeyDWORD(registryString, path, name, value)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "AddRegKeyDWORD").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmAddRegKeyExpandedString(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 4 {
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 4 {
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var registryString string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Errorf("Could not export field: %s", "registryString")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		registryString = rawArg0.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var path string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		path = filepath.Clean(rawArg1.(string))
	default:
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var name string
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Errorf("Could not export field: %s", "name")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case string:
		name = rawArg2.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var value string
	rawArg3, err := call.Argument(3).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Errorf("Could not export field: %s", "value")
		return otto.FalseValue()
	}
	switch v := rawArg3.(type) {
	case string:
		value = rawArg3.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	runtimeError := e.AddRegKeyExpandedString(registryString, path, name, value)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "AddRegKeyExpandedString").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmAddRegKeyQWORD(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 4 {
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 4 {
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var registryString string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Errorf("Could not export field: %s", "registryString")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		registryString = rawArg0.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var path string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		path = filepath.Clean(rawArg1.(string))
	default:
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var name string
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Errorf("Could not export field: %s", "name")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case string:
		name = rawArg2.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var value uint64
	rawArg3, err := call.Argument(3).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Errorf("Could not export field: %s", "value")
		return otto.FalseValue()
	}
	switch v := rawArg3.(type) {
	case uint64:
		value = rawArg3.(uint64)
	default:
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "uint64", v)
		return otto.FalseValue()
	}
	runtimeError := e.AddRegKeyQWORD(registryString, path, name, value)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "AddRegKeyQWORD").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmAddRegKeyString(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 4 {
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 4 {
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var registryString string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Errorf("Could not export field: %s", "registryString")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		registryString = rawArg0.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var path string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		path = filepath.Clean(rawArg1.(string))
	default:
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var name string
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Errorf("Could not export field: %s", "name")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case string:
		name = rawArg2.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var value string
	rawArg3, err := call.Argument(3).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Errorf("Could not export field: %s", "value")
		return otto.FalseValue()
	}
	switch v := rawArg3.(type) {
	case string:
		value = rawArg3.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	runtimeError := e.AddRegKeyString(registryString, path, name, value)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "AddRegKeyString").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmAddRegKeyStrings(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 4 {
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 4 {
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var registryString string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Errorf("Could not export field: %s", "registryString")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		registryString = rawArg0.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var path string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		path = filepath.Clean(rawArg1.(string))
	default:
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var name string
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Errorf("Could not export field: %s", "name")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case string:
		name = rawArg2.(string)
	default:
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var value []string
	rawArg3, err := call.Argument(3).Export()
	if err != nil {
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Errorf("Could not export field: %s", "value")
		return otto.FalseValue()
	}
	switch v := rawArg3.(type) {
	case []string:
		value = rawArg3.([]string)
	default:
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "[]string", v)
		return otto.FalseValue()
	}
	runtimeError := e.AddRegKeyStrings(registryString, path, name, value)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "AddRegKeyStrings").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmAppendFileBytes(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "AppendFileBytes").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "AppendFileBytes").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var path string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "AppendFileBytes").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		path = filepath.Clean(rawArg0.(string))
	default:
		e.Logger.WithField("function", "AppendFileBytes").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	fileData := e.ValueToByteSlice(call.Argument(1))
	fileError := e.AppendFileBytes(path, fileData)
	rawVMRet := VMResponse{}

	if fileError != nil {
		e.Logger.WithField("function", "AppendFileBytes").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "AppendFileBytes").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmAppendFileString(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "AppendFileString").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "AppendFileString").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var path string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "AppendFileString").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		path = filepath.Clean(rawArg0.(string))
	default:
		e.Logger.WithField("function", "AppendFileString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var addString string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "AppendFileString").WithField("trace", "true").Errorf("Could not export field: %s", "addString")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		addString = rawArg1.(string)
	default:
		e.Logger.WithField("function", "AppendFileString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	fileError := e.AppendFileString(path, addString)
	rawVMRet := VMResponse{}

	if fileError != nil {
		e.Logger.WithField("function", "AppendFileString").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "AppendFileString").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmAsset(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "Asset").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "Asset").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var assetName string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "Asset").WithField("trace", "true").Errorf("Could not export field: %s", "assetName")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		assetName = rawArg0.(string)
	default:
		e.Logger.WithField("function", "Asset").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	fileData, err := e.Asset(assetName)
	rawVMRet := VMResponse{}

	rawVMRet["fileData"] = fileData

	if err != nil {
		e.Logger.WithField("function", "Asset").WithField("trace", "true").Errorf("<function error> %s", err.Error())
		rawVMRet["err"] = err.Error()
	} else {
		rawVMRet["err"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "Asset").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmB64Decode(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "B64Decode").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "B64Decode").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var data string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "B64Decode").WithField("trace", "true").Errorf("Could not export field: %s", "data")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		data = rawArg0.(string)
	default:
		e.Logger.WithField("function", "B64Decode").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	value, execError := e.B64Decode(data)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value

	if execError != nil {
		e.Logger.WithField("function", "B64Decode").WithField("trace", "true").Errorf("<function error> %s", execError.Error())
		rawVMRet["execError"] = execError.Error()
	} else {
		rawVMRet["execError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "B64Decode").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmB64Encode(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "B64Encode").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "B64Encode").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	data := e.ValueToByteSlice(call.Argument(0))
	value := e.B64Encode(data)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "B64Encode").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmCheckIfCPUCountIsHigherThanOne(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "CheckIfCPUCountIsHigherThanOne").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "CheckIfCPUCountIsHigherThanOne").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	areWeInASandbox := e.CheckIfCPUCountIsHigherThanOne()
	rawVMRet := VMResponse{}

	rawVMRet["areWeInASandbox"] = areWeInASandbox
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "CheckIfCPUCountIsHigherThanOne").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmCheckIfRAMAmountIsBelow1GB(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "CheckIfRAMAmountIsBelow1GB").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "CheckIfRAMAmountIsBelow1GB").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	areWeInASandbox := e.CheckIfRAMAmountIsBelow1GB()
	rawVMRet := VMResponse{}

	rawVMRet["areWeInASandbox"] = areWeInASandbox
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "CheckIfRAMAmountIsBelow1GB").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmCheckIfWineGetUnixFileNameExists(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "CheckIfWineGetUnixFileNameExists").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "CheckIfWineGetUnixFileNameExists").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	areWeRunningInWine, runtimeError := e.CheckIfWineGetUnixFileNameExists()
	rawVMRet := VMResponse{}

	rawVMRet["areWeRunningInWine"] = areWeRunningInWine

	if runtimeError != nil {
		e.Logger.WithField("function", "CheckIfWineGetUnixFileNameExists").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "CheckIfWineGetUnixFileNameExists").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmCheckSandboxUsernames(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "CheckSandboxUsernames").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "CheckSandboxUsernames").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	areWeInASandbox, runtimeError := e.CheckSandboxUsernames()
	rawVMRet := VMResponse{}

	rawVMRet["areWeInASandbox"] = areWeInASandbox

	if runtimeError != nil {
		e.Logger.WithField("function", "CheckSandboxUsernames").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "CheckSandboxUsernames").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmChmod(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "Chmod").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "Chmod").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var path string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "Chmod").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		path = filepath.Clean(rawArg0.(string))
	default:
		e.Logger.WithField("function", "Chmod").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var perms int64
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "Chmod").WithField("trace", "true").Errorf("Could not export field: %s", "perms")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case int64:
		perms = rawArg1.(int64)
	default:
		e.Logger.WithField("function", "Chmod").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int64", v)
		return otto.FalseValue()
	}
	osError := e.Chmod(path, perms)
	rawVMRet := VMResponse{}

	if osError != nil {
		e.Logger.WithField("function", "Chmod").WithField("trace", "true").Errorf("<function error> %s", osError.Error())
		rawVMRet["osError"] = osError.Error()
	} else {
		rawVMRet["osError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "Chmod").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmCopyFile(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 3 {
		e.Logger.WithField("function", "CopyFile").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 3 {
		e.Logger.WithField("function", "CopyFile").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var srcPath string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "CopyFile").WithField("trace", "true").Errorf("Could not export field: %s", "srcPath")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		srcPath = rawArg0.(string)
	default:
		e.Logger.WithField("function", "CopyFile").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var dstPath string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "CopyFile").WithField("trace", "true").Errorf("Could not export field: %s", "dstPath")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		dstPath = rawArg1.(string)
	default:
		e.Logger.WithField("function", "CopyFile").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var perms int64
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "CopyFile").WithField("trace", "true").Errorf("Could not export field: %s", "perms")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case int64:
		perms = rawArg2.(int64)
	default:
		e.Logger.WithField("function", "CopyFile").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int64", v)
		return otto.FalseValue()
	}
	fileError := e.CopyFile(srcPath, dstPath, perms)
	rawVMRet := VMResponse{}

	if fileError != nil {
		e.Logger.WithField("function", "CopyFile").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "CopyFile").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmCreateDir(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "CreateDir").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "CreateDir").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var path string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "CreateDir").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		path = filepath.Clean(rawArg0.(string))
	default:
		e.Logger.WithField("function", "CreateDir").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	fileError := e.CreateDir(path)
	rawVMRet := VMResponse{}

	if fileError != nil {
		e.Logger.WithField("function", "CreateDir").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "CreateDir").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmDNSQuestion(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "DNSQuestion").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "DNSQuestion").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var target string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "DNSQuestion").WithField("trace", "true").Errorf("Could not export field: %s", "target")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		target = rawArg0.(string)
	default:
		e.Logger.WithField("function", "DNSQuestion").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var request string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "DNSQuestion").WithField("trace", "true").Errorf("Could not export field: %s", "request")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		request = rawArg1.(string)
	default:
		e.Logger.WithField("function", "DNSQuestion").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	answer, runtimeError := e.DNSQuestion(target, request)
	rawVMRet := VMResponse{}

	rawVMRet["answer"] = answer

	if runtimeError != nil {
		e.Logger.WithField("function", "DNSQuestion").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "DNSQuestion").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmDelRegKey(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "DelRegKey").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "DelRegKey").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var registryString string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "DelRegKey").WithField("trace", "true").Errorf("Could not export field: %s", "registryString")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		registryString = rawArg0.(string)
	default:
		e.Logger.WithField("function", "DelRegKey").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var path string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "DelRegKey").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		path = filepath.Clean(rawArg1.(string))
	default:
		e.Logger.WithField("function", "DelRegKey").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	runtimeError := e.DelRegKey(registryString, path)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "DelRegKey").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "DelRegKey").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmDelRegKeyValue(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 3 {
		e.Logger.WithField("function", "DelRegKeyValue").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 3 {
		e.Logger.WithField("function", "DelRegKeyValue").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var registryString string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "DelRegKeyValue").WithField("trace", "true").Errorf("Could not export field: %s", "registryString")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		registryString = rawArg0.(string)
	default:
		e.Logger.WithField("function", "DelRegKeyValue").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var path string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "DelRegKeyValue").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		path = filepath.Clean(rawArg1.(string))
	default:
		e.Logger.WithField("function", "DelRegKeyValue").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var value string
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "DelRegKeyValue").WithField("trace", "true").Errorf("Could not export field: %s", "value")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case string:
		value = rawArg2.(string)
	default:
		e.Logger.WithField("function", "DelRegKeyValue").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	runtimeError := e.DelRegKeyValue(registryString, path, value)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "DelRegKeyValue").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "DelRegKeyValue").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmDeleteFile(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "DeleteFile").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "DeleteFile").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var path string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "DeleteFile").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		path = filepath.Clean(rawArg0.(string))
	default:
		e.Logger.WithField("function", "DeleteFile").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	fileError := e.DeleteFile(path)
	rawVMRet := VMResponse{}

	if fileError != nil {
		e.Logger.WithField("function", "DeleteFile").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "DeleteFile").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmDeobfuscateString(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "DeobfuscateString").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "DeobfuscateString").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var str string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "DeobfuscateString").WithField("trace", "true").Errorf("Could not export field: %s", "str")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		str = rawArg0.(string)
	default:
		e.Logger.WithField("function", "DeobfuscateString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	value := e.DeobfuscateString(str)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "DeobfuscateString").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmEnvVars(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "EnvVars").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "EnvVars").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	vars := e.EnvVars()
	rawVMRet := VMResponse{}

	rawVMRet["vars"] = vars
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "EnvVars").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmExecuteCommand(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "ExecuteCommand").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "ExecuteCommand").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var baseCmd string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "ExecuteCommand").WithField("trace", "true").Errorf("Could not export field: %s", "baseCmd")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		baseCmd = rawArg0.(string)
	default:
		e.Logger.WithField("function", "ExecuteCommand").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var cmdArgs []string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "ExecuteCommand").WithField("trace", "true").Errorf("Could not export field: %s", "cmdArgs")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case []string:
		cmdArgs = rawArg1.([]string)
	default:
		e.Logger.WithField("function", "ExecuteCommand").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "[]string", v)
		return otto.FalseValue()
	}
	retObject := e.ExecuteCommand(baseCmd, cmdArgs)
	rawVMRet := VMResponse{}

	rawVMRet["retObject"] = retObject
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "ExecuteCommand").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmFileExists(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "FileExists").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "FileExists").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var path string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "FileExists").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		path = filepath.Clean(rawArg0.(string))
	default:
		e.Logger.WithField("function", "FileExists").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	fileExists, fileError := e.FileExists(path)
	rawVMRet := VMResponse{}

	rawVMRet["fileExists"] = fileExists

	if fileError != nil {
		e.Logger.WithField("function", "FileExists").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "FileExists").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmFindProcByName(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "FindProcByName").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "FindProcByName").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var procName string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "FindProcByName").WithField("trace", "true").Errorf("Could not export field: %s", "procName")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		procName = rawArg0.(string)
	default:
		e.Logger.WithField("function", "FindProcByName").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	pid, procError := e.FindProcByName(procName)
	rawVMRet := VMResponse{}

	rawVMRet["pid"] = pid

	if procError != nil {
		e.Logger.WithField("function", "FindProcByName").WithField("trace", "true").Errorf("<function error> %s", procError.Error())
		rawVMRet["procError"] = procError.Error()
	} else {
		rawVMRet["procError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "FindProcByName").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmForkExecuteCommand(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "ForkExecuteCommand").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "ForkExecuteCommand").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var baseCmd string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "ForkExecuteCommand").WithField("trace", "true").Errorf("Could not export field: %s", "baseCmd")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		baseCmd = rawArg0.(string)
	default:
		e.Logger.WithField("function", "ForkExecuteCommand").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var cmdArgs []string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "ForkExecuteCommand").WithField("trace", "true").Errorf("Could not export field: %s", "cmdArgs")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case []string:
		cmdArgs = rawArg1.([]string)
	default:
		e.Logger.WithField("function", "ForkExecuteCommand").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "[]string", v)
		return otto.FalseValue()
	}
	pid, execError := e.ForkExecuteCommand(baseCmd, cmdArgs)
	rawVMRet := VMResponse{}

	rawVMRet["pid"] = pid

	if execError != nil {
		e.Logger.WithField("function", "ForkExecuteCommand").WithField("trace", "true").Errorf("<function error> %s", execError.Error())
		rawVMRet["execError"] = execError.Error()
	} else {
		rawVMRet["execError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "ForkExecuteCommand").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmGetEnvVar(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "GetEnvVar").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "GetEnvVar").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var vars string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "GetEnvVar").WithField("trace", "true").Errorf("Could not export field: %s", "vars")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		vars = rawArg0.(string)
	default:
		e.Logger.WithField("function", "GetEnvVar").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	value := e.GetEnvVar(vars)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "GetEnvVar").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmGetLocalIPs(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "GetLocalIPs").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "GetLocalIPs").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	addresses := e.GetLocalIPs()
	rawVMRet := VMResponse{}

	rawVMRet["addresses"] = addresses
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "GetLocalIPs").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmGetMACAddress(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "GetMACAddress").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "GetMACAddress").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	address, runtimeError := e.GetMACAddress()
	rawVMRet := VMResponse{}

	rawVMRet["address"] = address

	if runtimeError != nil {
		e.Logger.WithField("function", "GetMACAddress").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "GetMACAddress").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmGetProcName(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "GetProcName").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "GetProcName").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var pid int
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "GetProcName").WithField("trace", "true").Errorf("Could not export field: %s", "pid")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case int:
		pid = rawArg0.(int)
	default:
		e.Logger.WithField("function", "GetProcName").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int", v)
		return otto.FalseValue()
	}
	procName, runtimeError := e.GetProcName(pid)
	rawVMRet := VMResponse{}

	rawVMRet["procName"] = procName

	if runtimeError != nil {
		e.Logger.WithField("function", "GetProcName").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "GetProcName").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmHTTPGetFile(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "HTTPGetFile").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "HTTPGetFile").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var url string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "HTTPGetFile").WithField("trace", "true").Errorf("Could not export field: %s", "url")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		url = rawArg0.(string)
	default:
		e.Logger.WithField("function", "HTTPGetFile").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	statusCode, file, runtimeError := e.HTTPGetFile(url)
	rawVMRet := VMResponse{}

	rawVMRet["statusCode"] = statusCode

	rawVMRet["file"] = file

	if runtimeError != nil {
		e.Logger.WithField("function", "HTTPGetFile").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "HTTPGetFile").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmHalt(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "Halt").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "Halt").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	value := e.Halt()
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "Halt").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmInjectIntoProc(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "InjectIntoProc").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "InjectIntoProc").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var shellcode string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "InjectIntoProc").WithField("trace", "true").Errorf("Could not export field: %s", "shellcode")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		shellcode = rawArg0.(string)
	default:
		e.Logger.WithField("function", "InjectIntoProc").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var proccessID int64
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "InjectIntoProc").WithField("trace", "true").Errorf("Could not export field: %s", "proccessID")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case int64:
		proccessID = rawArg1.(int64)
	default:
		e.Logger.WithField("function", "InjectIntoProc").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int64", v)
		return otto.FalseValue()
	}
	runtimeError := e.InjectIntoProc(shellcode, proccessID)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "InjectIntoProc").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "InjectIntoProc").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmInjectIntoSelf(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "InjectIntoSelf").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "InjectIntoSelf").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var shellcode string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "InjectIntoSelf").WithField("trace", "true").Errorf("Could not export field: %s", "shellcode")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		shellcode = rawArg0.(string)
	default:
		e.Logger.WithField("function", "InjectIntoSelf").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	runtimeError := e.InjectIntoSelf(shellcode)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "InjectIntoSelf").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "InjectIntoSelf").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmInstallSystemService(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 4 {
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 4 {
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var path string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		path = filepath.Clean(rawArg0.(string))
	default:
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var name string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Errorf("Could not export field: %s", "name")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		name = rawArg1.(string)
	default:
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var displayName string
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Errorf("Could not export field: %s", "displayName")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case string:
		displayName = rawArg2.(string)
	default:
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var description string
	rawArg3, err := call.Argument(3).Export()
	if err != nil {
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Errorf("Could not export field: %s", "description")
		return otto.FalseValue()
	}
	switch v := rawArg3.(type) {
	case string:
		description = rawArg3.(string)
	default:
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	installError := e.InstallSystemService(path, name, displayName, description)
	rawVMRet := VMResponse{}

	if installError != nil {
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Errorf("<function error> %s", installError.Error())
		rawVMRet["installError"] = installError.Error()
	} else {
		rawVMRet["installError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "InstallSystemService").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmIsDebuggerPresent(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "IsDebuggerPresent").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "IsDebuggerPresent").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	isBeingDebugged, runtimeError := e.IsDebuggerPresent()
	rawVMRet := VMResponse{}

	rawVMRet["isBeingDebugged"] = isBeingDebugged

	if runtimeError != nil {
		e.Logger.WithField("function", "IsDebuggerPresent").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "IsDebuggerPresent").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmIsTCPPortInUse(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "IsTCPPortInUse").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "IsTCPPortInUse").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var port string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "IsTCPPortInUse").WithField("trace", "true").Errorf("Could not export field: %s", "port")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		port = rawArg0.(string)
	default:
		e.Logger.WithField("function", "IsTCPPortInUse").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	state := e.IsTCPPortInUse(port)
	rawVMRet := VMResponse{}

	rawVMRet["state"] = state
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "IsTCPPortInUse").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmIsUDPPortInUse(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "IsUDPPortInUse").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "IsUDPPortInUse").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var port string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "IsUDPPortInUse").WithField("trace", "true").Errorf("Could not export field: %s", "port")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		port = rawArg0.(string)
	default:
		e.Logger.WithField("function", "IsUDPPortInUse").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	state := e.IsUDPPortInUse(port)
	rawVMRet := VMResponse{}

	rawVMRet["state"] = state
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "IsUDPPortInUse").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmMD5(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "MD5").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "MD5").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	data := e.ValueToByteSlice(call.Argument(0))
	value := e.MD5(data)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "MD5").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmMakeDebuggable(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "MakeDebuggable").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "MakeDebuggable").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	worked, runtimeError := e.MakeDebuggable()
	rawVMRet := VMResponse{}

	rawVMRet["worked"] = worked

	if runtimeError != nil {
		e.Logger.WithField("function", "MakeDebuggable").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "MakeDebuggable").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmMakeUnDebuggable(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "MakeUnDebuggable").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "MakeUnDebuggable").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	worked, runtimeError := e.MakeUnDebuggable()
	rawVMRet := VMResponse{}

	rawVMRet["worked"] = worked

	if runtimeError != nil {
		e.Logger.WithField("function", "MakeUnDebuggable").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "MakeUnDebuggable").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmModTime(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "ModTime").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "ModTime").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var path string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "ModTime").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		path = filepath.Clean(rawArg0.(string))
	default:
		e.Logger.WithField("function", "ModTime").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	modTime, fileError := e.ModTime(path)
	rawVMRet := VMResponse{}

	rawVMRet["modTime"] = modTime

	if fileError != nil {
		e.Logger.WithField("function", "ModTime").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "ModTime").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmModifyTimestamp(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 3 {
		e.Logger.WithField("function", "ModifyTimestamp").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 3 {
		e.Logger.WithField("function", "ModifyTimestamp").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var path string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "ModifyTimestamp").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		path = filepath.Clean(rawArg0.(string))
	default:
		e.Logger.WithField("function", "ModifyTimestamp").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var accessTime int64
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "ModifyTimestamp").WithField("trace", "true").Errorf("Could not export field: %s", "accessTime")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case int64:
		accessTime = rawArg1.(int64)
	default:
		e.Logger.WithField("function", "ModifyTimestamp").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int64", v)
		return otto.FalseValue()
	}

	var modifyTime int64
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "ModifyTimestamp").WithField("trace", "true").Errorf("Could not export field: %s", "modifyTime")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case int64:
		modifyTime = rawArg2.(int64)
	default:
		e.Logger.WithField("function", "ModifyTimestamp").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int64", v)
		return otto.FalseValue()
	}
	fileError := e.ModifyTimestamp(path, accessTime, modifyTime)
	rawVMRet := VMResponse{}

	if fileError != nil {
		e.Logger.WithField("function", "ModifyTimestamp").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "ModifyTimestamp").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmObfuscateString(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "ObfuscateString").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "ObfuscateString").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var str string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "ObfuscateString").WithField("trace", "true").Errorf("Could not export field: %s", "str")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		str = rawArg0.(string)
	default:
		e.Logger.WithField("function", "ObfuscateString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	value := e.ObfuscateString(str)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "ObfuscateString").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmPostJSON(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "PostJSON").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "PostJSON").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var url string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "PostJSON").WithField("trace", "true").Errorf("Could not export field: %s", "url")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		url = rawArg0.(string)
	default:
		e.Logger.WithField("function", "PostJSON").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var json string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "PostJSON").WithField("trace", "true").Errorf("Could not export field: %s", "json")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		json = rawArg1.(string)
	default:
		e.Logger.WithField("function", "PostJSON").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	statusCode, response, runtimeError := e.PostJSON(url, json)
	rawVMRet := VMResponse{}

	rawVMRet["statusCode"] = statusCode

	rawVMRet["response"] = response

	if runtimeError != nil {
		e.Logger.WithField("function", "PostJSON").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "PostJSON").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmQueryRegKey(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "QueryRegKey").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "QueryRegKey").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var registryString string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "QueryRegKey").WithField("trace", "true").Errorf("Could not export field: %s", "registryString")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		registryString = rawArg0.(string)
	default:
		e.Logger.WithField("function", "QueryRegKey").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var path string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "QueryRegKey").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		path = filepath.Clean(rawArg1.(string))
	default:
		e.Logger.WithField("function", "QueryRegKey").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	keyObj, runtimeError := e.QueryRegKey(registryString, path)
	rawVMRet := VMResponse{}

	rawVMRet["keyObj"] = keyObj

	if runtimeError != nil {
		e.Logger.WithField("function", "QueryRegKey").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "QueryRegKey").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmRandomInt(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "RandomInt").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "RandomInt").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var min int64
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "RandomInt").WithField("trace", "true").Errorf("Could not export field: %s", "min")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case int64:
		min = rawArg0.(int64)
	default:
		e.Logger.WithField("function", "RandomInt").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int64", v)
		return otto.FalseValue()
	}

	var max int64
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "RandomInt").WithField("trace", "true").Errorf("Could not export field: %s", "max")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case int64:
		max = rawArg1.(int64)
	default:
		e.Logger.WithField("function", "RandomInt").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int64", v)
		return otto.FalseValue()
	}
	value := e.RandomInt(min, max)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "RandomInt").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmRandomMixedCaseString(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "RandomMixedCaseString").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "RandomMixedCaseString").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var strlen int64
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "RandomMixedCaseString").WithField("trace", "true").Errorf("Could not export field: %s", "strlen")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case int64:
		strlen = rawArg0.(int64)
	default:
		e.Logger.WithField("function", "RandomMixedCaseString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int64", v)
		return otto.FalseValue()
	}
	value := e.RandomMixedCaseString(strlen)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "RandomMixedCaseString").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmRandomString(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "RandomString").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "RandomString").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var strlen int64
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "RandomString").WithField("trace", "true").Errorf("Could not export field: %s", "strlen")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case int64:
		strlen = rawArg0.(int64)
	default:
		e.Logger.WithField("function", "RandomString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int64", v)
		return otto.FalseValue()
	}
	value := e.RandomString(strlen)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "RandomString").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmReadFile(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "ReadFile").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "ReadFile").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var path string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "ReadFile").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		path = filepath.Clean(rawArg0.(string))
	default:
		e.Logger.WithField("function", "ReadFile").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	fileBytes, fileError := e.ReadFile(path)
	rawVMRet := VMResponse{}

	rawVMRet["fileBytes"] = fileBytes

	if fileError != nil {
		e.Logger.WithField("function", "ReadFile").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "ReadFile").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmRemoveServiceByName(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "RemoveServiceByName").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "RemoveServiceByName").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var name string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "RemoveServiceByName").WithField("trace", "true").Errorf("Could not export field: %s", "name")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		name = rawArg0.(string)
	default:
		e.Logger.WithField("function", "RemoveServiceByName").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	removalError := e.RemoveServiceByName(name)
	rawVMRet := VMResponse{}

	if removalError != nil {
		e.Logger.WithField("function", "RemoveServiceByName").WithField("trace", "true").Errorf("<function error> %s", removalError.Error())
		rawVMRet["removalError"] = removalError.Error()
	} else {
		rawVMRet["removalError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "RemoveServiceByName").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmReplaceFileString(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 3 {
		e.Logger.WithField("function", "ReplaceFileString").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 3 {
		e.Logger.WithField("function", "ReplaceFileString").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var file string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "ReplaceFileString").WithField("trace", "true").Errorf("Could not export field: %s", "file")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		file = rawArg0.(string)
	default:
		e.Logger.WithField("function", "ReplaceFileString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var match string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "ReplaceFileString").WithField("trace", "true").Errorf("Could not export field: %s", "match")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		match = rawArg1.(string)
	default:
		e.Logger.WithField("function", "ReplaceFileString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var replacement string
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "ReplaceFileString").WithField("trace", "true").Errorf("Could not export field: %s", "replacement")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case string:
		replacement = rawArg2.(string)
	default:
		e.Logger.WithField("function", "ReplaceFileString").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	stringsReplaced, fileError := e.ReplaceFileString(file, match, replacement)
	rawVMRet := VMResponse{}

	rawVMRet["stringsReplaced"] = stringsReplaced

	if fileError != nil {
		e.Logger.WithField("function", "ReplaceFileString").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "ReplaceFileString").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmRetrievePEPolymorphicData(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "RetrievePEPolymorphicData").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "RetrievePEPolymorphicData").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var peFile string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "RetrievePEPolymorphicData").WithField("trace", "true").Errorf("Could not export field: %s", "peFile")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		peFile = rawArg0.(string)
	default:
		e.Logger.WithField("function", "RetrievePEPolymorphicData").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	data, runtimeError := e.RetrievePEPolymorphicData(peFile)
	rawVMRet := VMResponse{}

	rawVMRet["data"] = data

	if runtimeError != nil {
		e.Logger.WithField("function", "RetrievePEPolymorphicData").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "RetrievePEPolymorphicData").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmRunningProcs(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "RunningProcs").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "RunningProcs").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	pids, runtimeError := e.RunningProcs()
	rawVMRet := VMResponse{}

	rawVMRet["pids"] = pids

	if runtimeError != nil {
		e.Logger.WithField("function", "RunningProcs").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "RunningProcs").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmSHA1(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "SHA1").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "SHA1").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	data := e.ValueToByteSlice(call.Argument(0))
	value := e.SHA1(data)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "SHA1").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmSSHCmd(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 5 {
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 5 {
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var hostAndPort string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Errorf("Could not export field: %s", "hostAndPort")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		hostAndPort = rawArg0.(string)
	default:
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var cmd string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Errorf("Could not export field: %s", "cmd")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		cmd = rawArg1.(string)
	default:
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var username string
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Errorf("Could not export field: %s", "username")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case string:
		username = rawArg2.(string)
	default:
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var password string
	rawArg3, err := call.Argument(3).Export()
	if err != nil {
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Errorf("Could not export field: %s", "password")
		return otto.FalseValue()
	}
	switch v := rawArg3.(type) {
	case string:
		password = rawArg3.(string)
	default:
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	key := e.ValueToByteSlice(call.Argument(4))
	response, runtimeError := e.SSHCmd(hostAndPort, cmd, username, password, key)
	rawVMRet := VMResponse{}

	rawVMRet["response"] = response

	if runtimeError != nil {
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "SSHCmd").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmSelfPath(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "SelfPath").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "SelfPath").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	path, osError := e.SelfPath()
	rawVMRet := VMResponse{}

	rawVMRet["path"] = path

	if osError != nil {
		e.Logger.WithField("function", "SelfPath").WithField("trace", "true").Errorf("<function error> %s", osError.Error())
		rawVMRet["osError"] = osError.Error()
	} else {
		rawVMRet["osError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "SelfPath").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmServePathOverHTTPS(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 3 {
		e.Logger.WithField("function", "ServePathOverHTTPS").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 3 {
		e.Logger.WithField("function", "ServePathOverHTTPS").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var port string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "ServePathOverHTTPS").WithField("trace", "true").Errorf("Could not export field: %s", "port")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		port = rawArg0.(string)
	default:
		e.Logger.WithField("function", "ServePathOverHTTPS").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var path string
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "ServePathOverHTTPS").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case string:
		path = filepath.Clean(rawArg1.(string))
	default:
		e.Logger.WithField("function", "ServePathOverHTTPS").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	var timeout int64
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "ServePathOverHTTPS").WithField("trace", "true").Errorf("Could not export field: %s", "timeout")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case int64:
		timeout = rawArg2.(int64)
	default:
		e.Logger.WithField("function", "ServePathOverHTTPS").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int64", v)
		return otto.FalseValue()
	}
	runtimeError := e.ServePathOverHTTPS(port, path, timeout)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "ServePathOverHTTPS").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "ServePathOverHTTPS").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmSignal(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "Signal").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "Signal").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var signal int
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "Signal").WithField("trace", "true").Errorf("Could not export field: %s", "signal")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case int:
		signal = rawArg0.(int)
	default:
		e.Logger.WithField("function", "Signal").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int", v)
		return otto.FalseValue()
	}

	var pid int
	rawArg1, err := call.Argument(1).Export()
	if err != nil {
		e.Logger.WithField("function", "Signal").WithField("trace", "true").Errorf("Could not export field: %s", "pid")
		return otto.FalseValue()
	}
	switch v := rawArg1.(type) {
	case int:
		pid = rawArg1.(int)
	default:
		e.Logger.WithField("function", "Signal").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int", v)
		return otto.FalseValue()
	}
	runtimeError := e.Signal(signal, pid)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "Signal").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "Signal").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmStartServiceByName(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "StartServiceByName").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "StartServiceByName").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var name string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "StartServiceByName").WithField("trace", "true").Errorf("Could not export field: %s", "name")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		name = rawArg0.(string)
	default:
		e.Logger.WithField("function", "StartServiceByName").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	startError := e.StartServiceByName(name)
	rawVMRet := VMResponse{}

	if startError != nil {
		e.Logger.WithField("function", "StartServiceByName").WithField("trace", "true").Errorf("<function error> %s", startError.Error())
		rawVMRet["startError"] = startError.Error()
	} else {
		rawVMRet["startError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "StartServiceByName").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmStopServiceByName(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "StopServiceByName").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "StopServiceByName").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var name string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "StopServiceByName").WithField("trace", "true").Errorf("Could not export field: %s", "name")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		name = rawArg0.(string)
	default:
		e.Logger.WithField("function", "StopServiceByName").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	installError := e.StopServiceByName(name)
	rawVMRet := VMResponse{}

	if installError != nil {
		e.Logger.WithField("function", "StopServiceByName").WithField("trace", "true").Errorf("<function error> %s", installError.Error())
		rawVMRet["installError"] = installError.Error()
	} else {
		rawVMRet["installError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "StopServiceByName").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmStripSpaces(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 1 {
		e.Logger.WithField("function", "StripSpaces").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 1 {
		e.Logger.WithField("function", "StripSpaces").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var str string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "StripSpaces").WithField("trace", "true").Errorf("Could not export field: %s", "str")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		str = rawArg0.(string)
	default:
		e.Logger.WithField("function", "StripSpaces").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}
	value := e.StripSpaces(str)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "StripSpaces").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmTimestamp(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		e.Logger.WithField("function", "Timestamp").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 0 {
		e.Logger.WithField("function", "Timestamp").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}
	value := e.Timestamp()
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "Timestamp").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmWriteFile(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 3 {
		e.Logger.WithField("function", "WriteFile").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 3 {
		e.Logger.WithField("function", "WriteFile").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var path string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "WriteFile").WithField("trace", "true").Errorf("Could not export field: %s", "path")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		path = filepath.Clean(rawArg0.(string))
	default:
		e.Logger.WithField("function", "WriteFile").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	fileData := e.ValueToByteSlice(call.Argument(1))

	var perms int64
	rawArg2, err := call.Argument(2).Export()
	if err != nil {
		e.Logger.WithField("function", "WriteFile").WithField("trace", "true").Errorf("Could not export field: %s", "perms")
		return otto.FalseValue()
	}
	switch v := rawArg2.(type) {
	case int64:
		perms = rawArg2.(int64)
	default:
		e.Logger.WithField("function", "WriteFile").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "int64", v)
		return otto.FalseValue()
	}
	bytesWritten, fileError := e.WriteFile(path, fileData, perms)
	rawVMRet := VMResponse{}

	rawVMRet["bytesWritten"] = bytesWritten

	if fileError != nil {
		e.Logger.WithField("function", "WriteFile").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "WriteFile").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmWritePEPolymorphicData(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "WritePEPolymorphicData").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "WritePEPolymorphicData").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var peFile string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "WritePEPolymorphicData").WithField("trace", "true").Errorf("Could not export field: %s", "peFile")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		peFile = rawArg0.(string)
	default:
		e.Logger.WithField("function", "WritePEPolymorphicData").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	data := e.ValueToByteSlice(call.Argument(1))
	runtimeError := e.WritePEPolymorphicData(peFile, data)
	rawVMRet := VMResponse{}

	if runtimeError != nil {
		e.Logger.WithField("function", "WritePEPolymorphicData").WithField("trace", "true").Errorf("<function error> %s", runtimeError.Error())
		rawVMRet["runtimeError"] = runtimeError.Error()
	} else {
		rawVMRet["runtimeError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "WritePEPolymorphicData").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmWriteTempFile(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "WriteTempFile").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "WriteTempFile").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	var name string
	rawArg0, err := call.Argument(0).Export()
	if err != nil {
		e.Logger.WithField("function", "WriteTempFile").WithField("trace", "true").Errorf("Could not export field: %s", "name")
		return otto.FalseValue()
	}
	switch v := rawArg0.(type) {
	case string:
		name = rawArg0.(string)
	default:
		e.Logger.WithField("function", "WriteTempFile").WithField("trace", "true").Errorf("Argument type mismatch: expected %s, got %T", "string", v)
		return otto.FalseValue()
	}

	fileData := e.ValueToByteSlice(call.Argument(1))
	fullpath, fileError := e.WriteTempFile(name, fileData)
	rawVMRet := VMResponse{}

	rawVMRet["fullpath"] = fullpath

	if fileError != nil {
		e.Logger.WithField("function", "WriteTempFile").WithField("trace", "true").Errorf("<function error> %s", fileError.Error())
		rawVMRet["fileError"] = fileError.Error()
	} else {
		rawVMRet["fileError"] = nil
	}
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "WriteTempFile").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}

func (e *Engine) vmXorBytes(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 2 {
		e.Logger.WithField("function", "XorBytes").WithField("trace", "true").Error("Too many arguments in call.")
		return otto.FalseValue()
	}
	if len(call.ArgumentList) < 2 {
		e.Logger.WithField("function", "XorBytes").WithField("trace", "true").Error("Too few arguments in call.")
		return otto.FalseValue()
	}

	aByteArray := e.ValueToByteSlice(call.Argument(0))

	bByteArray := e.ValueToByteSlice(call.Argument(1))
	value := e.XorBytes(aByteArray, bByteArray)
	rawVMRet := VMResponse{}

	rawVMRet["value"] = value
	vmRet, vmRetError := e.VM.ToValue(rawVMRet)
	if vmRetError != nil {
		e.Logger.WithField("function", "XorBytes").WithField("trace", "true").Errorf("Return conversion failed: %s", vmRetError.Error())
		return otto.FalseValue()
	}
	return vmRet
}
