package webkitrtcpeerconnection

import (
	"github.com/matthewmueller/golly/dom/domerror"
	"github.com/matthewmueller/golly/dom/mediastreamevent"
	"github.com/matthewmueller/golly/dom/rtcconfiguration"
	"github.com/matthewmueller/golly/dom/rtcicecandidate"
	"github.com/matthewmueller/golly/dom/rtciceconnectionstate"
	"github.com/matthewmueller/golly/dom/rtcicegatheringstate"
	"github.com/matthewmueller/golly/dom/rtcofferoptions"
	"github.com/matthewmueller/golly/dom/rtcpeerconnection"
	"github.com/matthewmueller/golly/dom/rtcpeerconnectioniceevent"
	"github.com/matthewmueller/golly/dom/rtcsessiondescription"
	"github.com/matthewmueller/golly/dom/rtcsignalingstate"
	"github.com/matthewmueller/golly/dom/rtcstatsreport"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ rtcpeerconnection.RTCPeerConnection = (*WebkitRTCPeerConnection)(nil)
var _ window.EventTarget = (*WebkitRTCPeerConnection)(nil)

// New fn
func New(configuration *rtcconfiguration.RTCConfiguration) *WebkitRTCPeerConnection {
	js.Rewrite("webkitRTCPeerConnection")
	return &WebkitRTCPeerConnection{}
}

// WebkitRTCPeerConnection struct
// js:"WebkitRTCPeerConnection,omit"
type WebkitRTCPeerConnection struct {
}

// AddIceCandidate fn
// js:"addIceCandidate"
func (*WebkitRTCPeerConnection) AddIceCandidate(candidate *rtcicecandidate.RTCIceCandidate, successCallback *func(), failureCallback *func(err *domerror.DOMError)) {
	js.Rewrite("await $_.addIceCandidate($1, $2, $3)", candidate, successCallback, failureCallback)
}

// AddStream fn
// js:"addStream"
func (*WebkitRTCPeerConnection) AddStream(stream *window.MediaStream) {
	js.Rewrite("$_.addStream($1)", stream)
}

// Close fn
// js:"close"
func (*WebkitRTCPeerConnection) Close() {
	js.Rewrite("$_.close()")
}

// CreateAnswer fn
// js:"createAnswer"
func (*WebkitRTCPeerConnection) CreateAnswer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError)) (r *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("await $_.createAnswer($1, $2)", successCallback, failureCallback)
	return r
}

// CreateOffer fn
// js:"createOffer"
func (*WebkitRTCPeerConnection) CreateOffer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError), options *rtcofferoptions.RTCOfferOptions) (r *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("await $_.createOffer($1, $2, $3)", successCallback, failureCallback, options)
	return r
}

// GetConfiguration fn
// js:"getConfiguration"
func (*WebkitRTCPeerConnection) GetConfiguration() (r *rtcconfiguration.RTCConfiguration) {
	js.Rewrite("$_.getConfiguration()")
	return r
}

// GetLocalStreams fn
// js:"getLocalStreams"
func (*WebkitRTCPeerConnection) GetLocalStreams() (w []*window.MediaStream) {
	js.Rewrite("$_.getLocalStreams()")
	return w
}

// GetRemoteStreams fn
// js:"getRemoteStreams"
func (*WebkitRTCPeerConnection) GetRemoteStreams() (w []*window.MediaStream) {
	js.Rewrite("$_.getRemoteStreams()")
	return w
}

// GetStats fn
// js:"getStats"
func (*WebkitRTCPeerConnection) GetStats(selector *window.MediaStreamTrack, successCallback *func(report *rtcstatsreport.RTCStatsReport), failureCallback *func(err *domerror.DOMError)) (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $_.getStats($1, $2, $3)", selector, successCallback, failureCallback)
	return r
}

// GetStreamByID fn
// js:"getStreamById"
func (*WebkitRTCPeerConnection) GetStreamByID(streamId string) (w *window.MediaStream) {
	js.Rewrite("$_.getStreamById($1)", streamId)
	return w
}

// RemoveStream fn
// js:"removeStream"
func (*WebkitRTCPeerConnection) RemoveStream(stream *window.MediaStream) {
	js.Rewrite("$_.removeStream($1)", stream)
}

// SetLocalDescription fn
// js:"setLocalDescription"
func (*WebkitRTCPeerConnection) SetLocalDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError)) {
	js.Rewrite("await $_.setLocalDescription($1, $2, $3)", description, successCallback, failureCallback)
}

// SetRemoteDescription fn
// js:"setRemoteDescription"
func (*WebkitRTCPeerConnection) SetRemoteDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError)) {
	js.Rewrite("await $_.setRemoteDescription($1, $2, $3)", description, successCallback, failureCallback)
}

// AddEventListener fn
// js:"addEventListener"
func (*WebkitRTCPeerConnection) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*WebkitRTCPeerConnection) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*WebkitRTCPeerConnection) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// CanTrickleIceCandidates prop
// js:"canTrickleIceCandidates"
func (*WebkitRTCPeerConnection) CanTrickleIceCandidates() (canTrickleIceCandidates bool) {
	js.Rewrite("$_.canTrickleIceCandidates")
	return canTrickleIceCandidates
}

// IceConnectionState prop
// js:"iceConnectionState"
func (*WebkitRTCPeerConnection) IceConnectionState() (iceConnectionState *rtciceconnectionstate.RTCIceConnectionState) {
	js.Rewrite("$_.iceConnectionState")
	return iceConnectionState
}

// IceGatheringState prop
// js:"iceGatheringState"
func (*WebkitRTCPeerConnection) IceGatheringState() (iceGatheringState *rtcicegatheringstate.RTCIceGatheringState) {
	js.Rewrite("$_.iceGatheringState")
	return iceGatheringState
}

// LocalDescription prop
// js:"localDescription"
func (*WebkitRTCPeerConnection) LocalDescription() (localDescription *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("$_.localDescription")
	return localDescription
}

// Onaddstream prop
// js:"onaddstream"
func (*WebkitRTCPeerConnection) Onaddstream() (onaddstream func(*mediastreamevent.MediaStreamEvent)) {
	js.Rewrite("$_.onaddstream")
	return onaddstream
}

// SetOnaddstream prop
// js:"onaddstream"
func (*WebkitRTCPeerConnection) SetOnaddstream(onaddstream func(*mediastreamevent.MediaStreamEvent)) {
	js.Rewrite("$_.onaddstream = $1", onaddstream)
}

// Onicecandidate prop
// js:"onicecandidate"
func (*WebkitRTCPeerConnection) Onicecandidate() (onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent)) {
	js.Rewrite("$_.onicecandidate")
	return onicecandidate
}

// SetOnicecandidate prop
// js:"onicecandidate"
func (*WebkitRTCPeerConnection) SetOnicecandidate(onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent)) {
	js.Rewrite("$_.onicecandidate = $1", onicecandidate)
}

// Oniceconnectionstatechange prop
// js:"oniceconnectionstatechange"
func (*WebkitRTCPeerConnection) Oniceconnectionstatechange() (oniceconnectionstatechange func(window.Event)) {
	js.Rewrite("$_.oniceconnectionstatechange")
	return oniceconnectionstatechange
}

// SetOniceconnectionstatechange prop
// js:"oniceconnectionstatechange"
func (*WebkitRTCPeerConnection) SetOniceconnectionstatechange(oniceconnectionstatechange func(window.Event)) {
	js.Rewrite("$_.oniceconnectionstatechange = $1", oniceconnectionstatechange)
}

// Onicegatheringstatechange prop
// js:"onicegatheringstatechange"
func (*WebkitRTCPeerConnection) Onicegatheringstatechange() (onicegatheringstatechange func(window.Event)) {
	js.Rewrite("$_.onicegatheringstatechange")
	return onicegatheringstatechange
}

// SetOnicegatheringstatechange prop
// js:"onicegatheringstatechange"
func (*WebkitRTCPeerConnection) SetOnicegatheringstatechange(onicegatheringstatechange func(window.Event)) {
	js.Rewrite("$_.onicegatheringstatechange = $1", onicegatheringstatechange)
}

// Onnegotiationneeded prop
// js:"onnegotiationneeded"
func (*WebkitRTCPeerConnection) Onnegotiationneeded() (onnegotiationneeded func(window.Event)) {
	js.Rewrite("$_.onnegotiationneeded")
	return onnegotiationneeded
}

// SetOnnegotiationneeded prop
// js:"onnegotiationneeded"
func (*WebkitRTCPeerConnection) SetOnnegotiationneeded(onnegotiationneeded func(window.Event)) {
	js.Rewrite("$_.onnegotiationneeded = $1", onnegotiationneeded)
}

// Onremovestream prop
// js:"onremovestream"
func (*WebkitRTCPeerConnection) Onremovestream() (onremovestream func(*mediastreamevent.MediaStreamEvent)) {
	js.Rewrite("$_.onremovestream")
	return onremovestream
}

// SetOnremovestream prop
// js:"onremovestream"
func (*WebkitRTCPeerConnection) SetOnremovestream(onremovestream func(*mediastreamevent.MediaStreamEvent)) {
	js.Rewrite("$_.onremovestream = $1", onremovestream)
}

// Onsignalingstatechange prop
// js:"onsignalingstatechange"
func (*WebkitRTCPeerConnection) Onsignalingstatechange() (onsignalingstatechange func(window.Event)) {
	js.Rewrite("$_.onsignalingstatechange")
	return onsignalingstatechange
}

// SetOnsignalingstatechange prop
// js:"onsignalingstatechange"
func (*WebkitRTCPeerConnection) SetOnsignalingstatechange(onsignalingstatechange func(window.Event)) {
	js.Rewrite("$_.onsignalingstatechange = $1", onsignalingstatechange)
}

// RemoteDescription prop
// js:"remoteDescription"
func (*WebkitRTCPeerConnection) RemoteDescription() (remoteDescription *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("$_.remoteDescription")
	return remoteDescription
}

// SignalingState prop
// js:"signalingState"
func (*WebkitRTCPeerConnection) SignalingState() (signalingState *rtcsignalingstate.RTCSignalingState) {
	js.Rewrite("$_.signalingState")
	return signalingState
}