package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	types.DefaultVMContext
}

func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{
		body: []byte("Hello!"),
	}
}

type pluginContext struct {
	types.DefaultPluginContext

	body []byte
}

func (ctx *pluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	data, err := proxywasm.GetPluginConfiguration()
	if err != nil {
		proxywasm.LogCriticalf("failed to read plugin configuration: %v", err)
	}
	proxywasm.LogDebugf("plugin configuration: %s", data)

	if data != nil {
		ctx.body = data
	}

	return types.OnPluginStartStatusOK
}

func (ctx *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &directResponseContext{
		body: ctx.body,
	}
}

type directResponseContext struct {
	types.DefaultHttpContext

	body []byte
}

func (ctx *directResponseContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	xFoo, err := proxywasm.GetHttpRequestHeader("x-foo")
	if err != nil && err != types.ErrorStatusNotFound {
		proxywasm.LogErrorf("failed to get x-foo header value: %v", err)
		return types.ActionContinue
	}

	xBar, err := proxywasm.GetHttpRequestHeader("x-bar")
	if err != nil && err != types.ErrorStatusNotFound {
		proxywasm.LogErrorf("failed to get x-bar header value: %v", err)
		return types.ActionContinue
	}

	proxywasm.LogDebugf("x-foo: %s, x-bar: %s", xFoo, xBar)
	if xFoo == "1" && xBar == "1" {
		return types.ActionContinue
	}

	if err := proxywasm.SendHttpResponse(400, nil, ctx.body, -1); err != nil {
		proxywasm.LogErrorf("failed to return direct response: %v", err)
		return types.ActionContinue
	}

	return types.ActionPause
}
