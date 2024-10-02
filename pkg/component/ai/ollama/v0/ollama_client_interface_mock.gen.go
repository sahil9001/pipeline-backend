// Code generated by http://github.com/gojuno/minimock (v3.4.0). DO NOT EDIT.

package ollama

//go:generate minimock -i github.com/instill-ai/pipeline-backend/pkg/component/ai/ollama/v0.OllamaClientInterface -o ollama_client_interface_mock.gen.go -n OllamaClientInterfaceMock -p ollama

import (
	_ "embed"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// OllamaClientInterfaceMock implements OllamaClientInterface
type OllamaClientInterfaceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcChat          func(c1 ChatRequest) (c2 ChatResponse, err error)
	funcChatOrigin    string
	inspectFuncChat   func(c1 ChatRequest)
	afterChatCounter  uint64
	beforeChatCounter uint64
	ChatMock          mOllamaClientInterfaceMockChat

	funcEmbed          func(e1 EmbedRequest) (e2 EmbedResponse, err error)
	funcEmbedOrigin    string
	inspectFuncEmbed   func(e1 EmbedRequest)
	afterEmbedCounter  uint64
	beforeEmbedCounter uint64
	EmbedMock          mOllamaClientInterfaceMockEmbed

	funcIsAutoPull          func() (b1 bool)
	funcIsAutoPullOrigin    string
	inspectFuncIsAutoPull   func()
	afterIsAutoPullCounter  uint64
	beforeIsAutoPullCounter uint64
	IsAutoPullMock          mOllamaClientInterfaceMockIsAutoPull
}

// NewOllamaClientInterfaceMock returns a mock for OllamaClientInterface
func NewOllamaClientInterfaceMock(t minimock.Tester) *OllamaClientInterfaceMock {
	m := &OllamaClientInterfaceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ChatMock = mOllamaClientInterfaceMockChat{mock: m}
	m.ChatMock.callArgs = []*OllamaClientInterfaceMockChatParams{}

	m.EmbedMock = mOllamaClientInterfaceMockEmbed{mock: m}
	m.EmbedMock.callArgs = []*OllamaClientInterfaceMockEmbedParams{}

	m.IsAutoPullMock = mOllamaClientInterfaceMockIsAutoPull{mock: m}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mOllamaClientInterfaceMockChat struct {
	optional           bool
	mock               *OllamaClientInterfaceMock
	defaultExpectation *OllamaClientInterfaceMockChatExpectation
	expectations       []*OllamaClientInterfaceMockChatExpectation

	callArgs []*OllamaClientInterfaceMockChatParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// OllamaClientInterfaceMockChatExpectation specifies expectation struct of the OllamaClientInterface.Chat
type OllamaClientInterfaceMockChatExpectation struct {
	mock               *OllamaClientInterfaceMock
	params             *OllamaClientInterfaceMockChatParams
	paramPtrs          *OllamaClientInterfaceMockChatParamPtrs
	expectationOrigins OllamaClientInterfaceMockChatExpectationOrigins
	results            *OllamaClientInterfaceMockChatResults
	returnOrigin       string
	Counter            uint64
}

// OllamaClientInterfaceMockChatParams contains parameters of the OllamaClientInterface.Chat
type OllamaClientInterfaceMockChatParams struct {
	c1 ChatRequest
}

// OllamaClientInterfaceMockChatParamPtrs contains pointers to parameters of the OllamaClientInterface.Chat
type OllamaClientInterfaceMockChatParamPtrs struct {
	c1 *ChatRequest
}

// OllamaClientInterfaceMockChatResults contains results of the OllamaClientInterface.Chat
type OllamaClientInterfaceMockChatResults struct {
	c2  ChatResponse
	err error
}

// OllamaClientInterfaceMockChatOrigins contains origins of expectations of the OllamaClientInterface.Chat
type OllamaClientInterfaceMockChatExpectationOrigins struct {
	origin   string
	originC1 string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmChat *mOllamaClientInterfaceMockChat) Optional() *mOllamaClientInterfaceMockChat {
	mmChat.optional = true
	return mmChat
}

// Expect sets up expected params for OllamaClientInterface.Chat
func (mmChat *mOllamaClientInterfaceMockChat) Expect(c1 ChatRequest) *mOllamaClientInterfaceMockChat {
	if mmChat.mock.funcChat != nil {
		mmChat.mock.t.Fatalf("OllamaClientInterfaceMock.Chat mock is already set by Set")
	}

	if mmChat.defaultExpectation == nil {
		mmChat.defaultExpectation = &OllamaClientInterfaceMockChatExpectation{}
	}

	if mmChat.defaultExpectation.paramPtrs != nil {
		mmChat.mock.t.Fatalf("OllamaClientInterfaceMock.Chat mock is already set by ExpectParams functions")
	}

	mmChat.defaultExpectation.params = &OllamaClientInterfaceMockChatParams{c1}
	mmChat.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmChat.expectations {
		if minimock.Equal(e.params, mmChat.defaultExpectation.params) {
			mmChat.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmChat.defaultExpectation.params)
		}
	}

	return mmChat
}

// ExpectC1Param1 sets up expected param c1 for OllamaClientInterface.Chat
func (mmChat *mOllamaClientInterfaceMockChat) ExpectC1Param1(c1 ChatRequest) *mOllamaClientInterfaceMockChat {
	if mmChat.mock.funcChat != nil {
		mmChat.mock.t.Fatalf("OllamaClientInterfaceMock.Chat mock is already set by Set")
	}

	if mmChat.defaultExpectation == nil {
		mmChat.defaultExpectation = &OllamaClientInterfaceMockChatExpectation{}
	}

	if mmChat.defaultExpectation.params != nil {
		mmChat.mock.t.Fatalf("OllamaClientInterfaceMock.Chat mock is already set by Expect")
	}

	if mmChat.defaultExpectation.paramPtrs == nil {
		mmChat.defaultExpectation.paramPtrs = &OllamaClientInterfaceMockChatParamPtrs{}
	}
	mmChat.defaultExpectation.paramPtrs.c1 = &c1
	mmChat.defaultExpectation.expectationOrigins.originC1 = minimock.CallerInfo(1)

	return mmChat
}

// Inspect accepts an inspector function that has same arguments as the OllamaClientInterface.Chat
func (mmChat *mOllamaClientInterfaceMockChat) Inspect(f func(c1 ChatRequest)) *mOllamaClientInterfaceMockChat {
	if mmChat.mock.inspectFuncChat != nil {
		mmChat.mock.t.Fatalf("Inspect function is already set for OllamaClientInterfaceMock.Chat")
	}

	mmChat.mock.inspectFuncChat = f

	return mmChat
}

// Return sets up results that will be returned by OllamaClientInterface.Chat
func (mmChat *mOllamaClientInterfaceMockChat) Return(c2 ChatResponse, err error) *OllamaClientInterfaceMock {
	if mmChat.mock.funcChat != nil {
		mmChat.mock.t.Fatalf("OllamaClientInterfaceMock.Chat mock is already set by Set")
	}

	if mmChat.defaultExpectation == nil {
		mmChat.defaultExpectation = &OllamaClientInterfaceMockChatExpectation{mock: mmChat.mock}
	}
	mmChat.defaultExpectation.results = &OllamaClientInterfaceMockChatResults{c2, err}
	mmChat.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmChat.mock
}

// Set uses given function f to mock the OllamaClientInterface.Chat method
func (mmChat *mOllamaClientInterfaceMockChat) Set(f func(c1 ChatRequest) (c2 ChatResponse, err error)) *OllamaClientInterfaceMock {
	if mmChat.defaultExpectation != nil {
		mmChat.mock.t.Fatalf("Default expectation is already set for the OllamaClientInterface.Chat method")
	}

	if len(mmChat.expectations) > 0 {
		mmChat.mock.t.Fatalf("Some expectations are already set for the OllamaClientInterface.Chat method")
	}

	mmChat.mock.funcChat = f
	mmChat.mock.funcChatOrigin = minimock.CallerInfo(1)
	return mmChat.mock
}

// When sets expectation for the OllamaClientInterface.Chat which will trigger the result defined by the following
// Then helper
func (mmChat *mOllamaClientInterfaceMockChat) When(c1 ChatRequest) *OllamaClientInterfaceMockChatExpectation {
	if mmChat.mock.funcChat != nil {
		mmChat.mock.t.Fatalf("OllamaClientInterfaceMock.Chat mock is already set by Set")
	}

	expectation := &OllamaClientInterfaceMockChatExpectation{
		mock:               mmChat.mock,
		params:             &OllamaClientInterfaceMockChatParams{c1},
		expectationOrigins: OllamaClientInterfaceMockChatExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmChat.expectations = append(mmChat.expectations, expectation)
	return expectation
}

// Then sets up OllamaClientInterface.Chat return parameters for the expectation previously defined by the When method
func (e *OllamaClientInterfaceMockChatExpectation) Then(c2 ChatResponse, err error) *OllamaClientInterfaceMock {
	e.results = &OllamaClientInterfaceMockChatResults{c2, err}
	return e.mock
}

// Times sets number of times OllamaClientInterface.Chat should be invoked
func (mmChat *mOllamaClientInterfaceMockChat) Times(n uint64) *mOllamaClientInterfaceMockChat {
	if n == 0 {
		mmChat.mock.t.Fatalf("Times of OllamaClientInterfaceMock.Chat mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmChat.expectedInvocations, n)
	mmChat.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmChat
}

func (mmChat *mOllamaClientInterfaceMockChat) invocationsDone() bool {
	if len(mmChat.expectations) == 0 && mmChat.defaultExpectation == nil && mmChat.mock.funcChat == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmChat.mock.afterChatCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmChat.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Chat implements OllamaClientInterface
func (mmChat *OllamaClientInterfaceMock) Chat(c1 ChatRequest) (c2 ChatResponse, err error) {
	mm_atomic.AddUint64(&mmChat.beforeChatCounter, 1)
	defer mm_atomic.AddUint64(&mmChat.afterChatCounter, 1)

	mmChat.t.Helper()

	if mmChat.inspectFuncChat != nil {
		mmChat.inspectFuncChat(c1)
	}

	mm_params := OllamaClientInterfaceMockChatParams{c1}

	// Record call args
	mmChat.ChatMock.mutex.Lock()
	mmChat.ChatMock.callArgs = append(mmChat.ChatMock.callArgs, &mm_params)
	mmChat.ChatMock.mutex.Unlock()

	for _, e := range mmChat.ChatMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.c2, e.results.err
		}
	}

	if mmChat.ChatMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmChat.ChatMock.defaultExpectation.Counter, 1)
		mm_want := mmChat.ChatMock.defaultExpectation.params
		mm_want_ptrs := mmChat.ChatMock.defaultExpectation.paramPtrs

		mm_got := OllamaClientInterfaceMockChatParams{c1}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.c1 != nil && !minimock.Equal(*mm_want_ptrs.c1, mm_got.c1) {
				mmChat.t.Errorf("OllamaClientInterfaceMock.Chat got unexpected parameter c1, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmChat.ChatMock.defaultExpectation.expectationOrigins.originC1, *mm_want_ptrs.c1, mm_got.c1, minimock.Diff(*mm_want_ptrs.c1, mm_got.c1))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmChat.t.Errorf("OllamaClientInterfaceMock.Chat got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmChat.ChatMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmChat.ChatMock.defaultExpectation.results
		if mm_results == nil {
			mmChat.t.Fatal("No results are set for the OllamaClientInterfaceMock.Chat")
		}
		return (*mm_results).c2, (*mm_results).err
	}
	if mmChat.funcChat != nil {
		return mmChat.funcChat(c1)
	}
	mmChat.t.Fatalf("Unexpected call to OllamaClientInterfaceMock.Chat. %v", c1)
	return
}

// ChatAfterCounter returns a count of finished OllamaClientInterfaceMock.Chat invocations
func (mmChat *OllamaClientInterfaceMock) ChatAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmChat.afterChatCounter)
}

// ChatBeforeCounter returns a count of OllamaClientInterfaceMock.Chat invocations
func (mmChat *OllamaClientInterfaceMock) ChatBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmChat.beforeChatCounter)
}

// Calls returns a list of arguments used in each call to OllamaClientInterfaceMock.Chat.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmChat *mOllamaClientInterfaceMockChat) Calls() []*OllamaClientInterfaceMockChatParams {
	mmChat.mutex.RLock()

	argCopy := make([]*OllamaClientInterfaceMockChatParams, len(mmChat.callArgs))
	copy(argCopy, mmChat.callArgs)

	mmChat.mutex.RUnlock()

	return argCopy
}

// MinimockChatDone returns true if the count of the Chat invocations corresponds
// the number of defined expectations
func (m *OllamaClientInterfaceMock) MinimockChatDone() bool {
	if m.ChatMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.ChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.ChatMock.invocationsDone()
}

// MinimockChatInspect logs each unmet expectation
func (m *OllamaClientInterfaceMock) MinimockChatInspect() {
	for _, e := range m.ChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OllamaClientInterfaceMock.Chat at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterChatCounter := mm_atomic.LoadUint64(&m.afterChatCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.ChatMock.defaultExpectation != nil && afterChatCounter < 1 {
		if m.ChatMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to OllamaClientInterfaceMock.Chat at\n%s", m.ChatMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to OllamaClientInterfaceMock.Chat at\n%s with params: %#v", m.ChatMock.defaultExpectation.expectationOrigins.origin, *m.ChatMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcChat != nil && afterChatCounter < 1 {
		m.t.Errorf("Expected call to OllamaClientInterfaceMock.Chat at\n%s", m.funcChatOrigin)
	}

	if !m.ChatMock.invocationsDone() && afterChatCounter > 0 {
		m.t.Errorf("Expected %d calls to OllamaClientInterfaceMock.Chat at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.ChatMock.expectedInvocations), m.ChatMock.expectedInvocationsOrigin, afterChatCounter)
	}
}

type mOllamaClientInterfaceMockEmbed struct {
	optional           bool
	mock               *OllamaClientInterfaceMock
	defaultExpectation *OllamaClientInterfaceMockEmbedExpectation
	expectations       []*OllamaClientInterfaceMockEmbedExpectation

	callArgs []*OllamaClientInterfaceMockEmbedParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// OllamaClientInterfaceMockEmbedExpectation specifies expectation struct of the OllamaClientInterface.Embed
type OllamaClientInterfaceMockEmbedExpectation struct {
	mock               *OllamaClientInterfaceMock
	params             *OllamaClientInterfaceMockEmbedParams
	paramPtrs          *OllamaClientInterfaceMockEmbedParamPtrs
	expectationOrigins OllamaClientInterfaceMockEmbedExpectationOrigins
	results            *OllamaClientInterfaceMockEmbedResults
	returnOrigin       string
	Counter            uint64
}

// OllamaClientInterfaceMockEmbedParams contains parameters of the OllamaClientInterface.Embed
type OllamaClientInterfaceMockEmbedParams struct {
	e1 EmbedRequest
}

// OllamaClientInterfaceMockEmbedParamPtrs contains pointers to parameters of the OllamaClientInterface.Embed
type OllamaClientInterfaceMockEmbedParamPtrs struct {
	e1 *EmbedRequest
}

// OllamaClientInterfaceMockEmbedResults contains results of the OllamaClientInterface.Embed
type OllamaClientInterfaceMockEmbedResults struct {
	e2  EmbedResponse
	err error
}

// OllamaClientInterfaceMockEmbedOrigins contains origins of expectations of the OllamaClientInterface.Embed
type OllamaClientInterfaceMockEmbedExpectationOrigins struct {
	origin   string
	originE1 string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmEmbed *mOllamaClientInterfaceMockEmbed) Optional() *mOllamaClientInterfaceMockEmbed {
	mmEmbed.optional = true
	return mmEmbed
}

// Expect sets up expected params for OllamaClientInterface.Embed
func (mmEmbed *mOllamaClientInterfaceMockEmbed) Expect(e1 EmbedRequest) *mOllamaClientInterfaceMockEmbed {
	if mmEmbed.mock.funcEmbed != nil {
		mmEmbed.mock.t.Fatalf("OllamaClientInterfaceMock.Embed mock is already set by Set")
	}

	if mmEmbed.defaultExpectation == nil {
		mmEmbed.defaultExpectation = &OllamaClientInterfaceMockEmbedExpectation{}
	}

	if mmEmbed.defaultExpectation.paramPtrs != nil {
		mmEmbed.mock.t.Fatalf("OllamaClientInterfaceMock.Embed mock is already set by ExpectParams functions")
	}

	mmEmbed.defaultExpectation.params = &OllamaClientInterfaceMockEmbedParams{e1}
	mmEmbed.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmEmbed.expectations {
		if minimock.Equal(e.params, mmEmbed.defaultExpectation.params) {
			mmEmbed.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmEmbed.defaultExpectation.params)
		}
	}

	return mmEmbed
}

// ExpectE1Param1 sets up expected param e1 for OllamaClientInterface.Embed
func (mmEmbed *mOllamaClientInterfaceMockEmbed) ExpectE1Param1(e1 EmbedRequest) *mOllamaClientInterfaceMockEmbed {
	if mmEmbed.mock.funcEmbed != nil {
		mmEmbed.mock.t.Fatalf("OllamaClientInterfaceMock.Embed mock is already set by Set")
	}

	if mmEmbed.defaultExpectation == nil {
		mmEmbed.defaultExpectation = &OllamaClientInterfaceMockEmbedExpectation{}
	}

	if mmEmbed.defaultExpectation.params != nil {
		mmEmbed.mock.t.Fatalf("OllamaClientInterfaceMock.Embed mock is already set by Expect")
	}

	if mmEmbed.defaultExpectation.paramPtrs == nil {
		mmEmbed.defaultExpectation.paramPtrs = &OllamaClientInterfaceMockEmbedParamPtrs{}
	}
	mmEmbed.defaultExpectation.paramPtrs.e1 = &e1
	mmEmbed.defaultExpectation.expectationOrigins.originE1 = minimock.CallerInfo(1)

	return mmEmbed
}

// Inspect accepts an inspector function that has same arguments as the OllamaClientInterface.Embed
func (mmEmbed *mOllamaClientInterfaceMockEmbed) Inspect(f func(e1 EmbedRequest)) *mOllamaClientInterfaceMockEmbed {
	if mmEmbed.mock.inspectFuncEmbed != nil {
		mmEmbed.mock.t.Fatalf("Inspect function is already set for OllamaClientInterfaceMock.Embed")
	}

	mmEmbed.mock.inspectFuncEmbed = f

	return mmEmbed
}

// Return sets up results that will be returned by OllamaClientInterface.Embed
func (mmEmbed *mOllamaClientInterfaceMockEmbed) Return(e2 EmbedResponse, err error) *OllamaClientInterfaceMock {
	if mmEmbed.mock.funcEmbed != nil {
		mmEmbed.mock.t.Fatalf("OllamaClientInterfaceMock.Embed mock is already set by Set")
	}

	if mmEmbed.defaultExpectation == nil {
		mmEmbed.defaultExpectation = &OllamaClientInterfaceMockEmbedExpectation{mock: mmEmbed.mock}
	}
	mmEmbed.defaultExpectation.results = &OllamaClientInterfaceMockEmbedResults{e2, err}
	mmEmbed.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmEmbed.mock
}

// Set uses given function f to mock the OllamaClientInterface.Embed method
func (mmEmbed *mOllamaClientInterfaceMockEmbed) Set(f func(e1 EmbedRequest) (e2 EmbedResponse, err error)) *OllamaClientInterfaceMock {
	if mmEmbed.defaultExpectation != nil {
		mmEmbed.mock.t.Fatalf("Default expectation is already set for the OllamaClientInterface.Embed method")
	}

	if len(mmEmbed.expectations) > 0 {
		mmEmbed.mock.t.Fatalf("Some expectations are already set for the OllamaClientInterface.Embed method")
	}

	mmEmbed.mock.funcEmbed = f
	mmEmbed.mock.funcEmbedOrigin = minimock.CallerInfo(1)
	return mmEmbed.mock
}

// When sets expectation for the OllamaClientInterface.Embed which will trigger the result defined by the following
// Then helper
func (mmEmbed *mOllamaClientInterfaceMockEmbed) When(e1 EmbedRequest) *OllamaClientInterfaceMockEmbedExpectation {
	if mmEmbed.mock.funcEmbed != nil {
		mmEmbed.mock.t.Fatalf("OllamaClientInterfaceMock.Embed mock is already set by Set")
	}

	expectation := &OllamaClientInterfaceMockEmbedExpectation{
		mock:               mmEmbed.mock,
		params:             &OllamaClientInterfaceMockEmbedParams{e1},
		expectationOrigins: OllamaClientInterfaceMockEmbedExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmEmbed.expectations = append(mmEmbed.expectations, expectation)
	return expectation
}

// Then sets up OllamaClientInterface.Embed return parameters for the expectation previously defined by the When method
func (e *OllamaClientInterfaceMockEmbedExpectation) Then(e2 EmbedResponse, err error) *OllamaClientInterfaceMock {
	e.results = &OllamaClientInterfaceMockEmbedResults{e2, err}
	return e.mock
}

// Times sets number of times OllamaClientInterface.Embed should be invoked
func (mmEmbed *mOllamaClientInterfaceMockEmbed) Times(n uint64) *mOllamaClientInterfaceMockEmbed {
	if n == 0 {
		mmEmbed.mock.t.Fatalf("Times of OllamaClientInterfaceMock.Embed mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmEmbed.expectedInvocations, n)
	mmEmbed.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmEmbed
}

func (mmEmbed *mOllamaClientInterfaceMockEmbed) invocationsDone() bool {
	if len(mmEmbed.expectations) == 0 && mmEmbed.defaultExpectation == nil && mmEmbed.mock.funcEmbed == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmEmbed.mock.afterEmbedCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmEmbed.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Embed implements OllamaClientInterface
func (mmEmbed *OllamaClientInterfaceMock) Embed(e1 EmbedRequest) (e2 EmbedResponse, err error) {
	mm_atomic.AddUint64(&mmEmbed.beforeEmbedCounter, 1)
	defer mm_atomic.AddUint64(&mmEmbed.afterEmbedCounter, 1)

	mmEmbed.t.Helper()

	if mmEmbed.inspectFuncEmbed != nil {
		mmEmbed.inspectFuncEmbed(e1)
	}

	mm_params := OllamaClientInterfaceMockEmbedParams{e1}

	// Record call args
	mmEmbed.EmbedMock.mutex.Lock()
	mmEmbed.EmbedMock.callArgs = append(mmEmbed.EmbedMock.callArgs, &mm_params)
	mmEmbed.EmbedMock.mutex.Unlock()

	for _, e := range mmEmbed.EmbedMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.e2, e.results.err
		}
	}

	if mmEmbed.EmbedMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmEmbed.EmbedMock.defaultExpectation.Counter, 1)
		mm_want := mmEmbed.EmbedMock.defaultExpectation.params
		mm_want_ptrs := mmEmbed.EmbedMock.defaultExpectation.paramPtrs

		mm_got := OllamaClientInterfaceMockEmbedParams{e1}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.e1 != nil && !minimock.Equal(*mm_want_ptrs.e1, mm_got.e1) {
				mmEmbed.t.Errorf("OllamaClientInterfaceMock.Embed got unexpected parameter e1, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmEmbed.EmbedMock.defaultExpectation.expectationOrigins.originE1, *mm_want_ptrs.e1, mm_got.e1, minimock.Diff(*mm_want_ptrs.e1, mm_got.e1))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmEmbed.t.Errorf("OllamaClientInterfaceMock.Embed got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmEmbed.EmbedMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmEmbed.EmbedMock.defaultExpectation.results
		if mm_results == nil {
			mmEmbed.t.Fatal("No results are set for the OllamaClientInterfaceMock.Embed")
		}
		return (*mm_results).e2, (*mm_results).err
	}
	if mmEmbed.funcEmbed != nil {
		return mmEmbed.funcEmbed(e1)
	}
	mmEmbed.t.Fatalf("Unexpected call to OllamaClientInterfaceMock.Embed. %v", e1)
	return
}

// EmbedAfterCounter returns a count of finished OllamaClientInterfaceMock.Embed invocations
func (mmEmbed *OllamaClientInterfaceMock) EmbedAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEmbed.afterEmbedCounter)
}

// EmbedBeforeCounter returns a count of OllamaClientInterfaceMock.Embed invocations
func (mmEmbed *OllamaClientInterfaceMock) EmbedBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEmbed.beforeEmbedCounter)
}

// Calls returns a list of arguments used in each call to OllamaClientInterfaceMock.Embed.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmEmbed *mOllamaClientInterfaceMockEmbed) Calls() []*OllamaClientInterfaceMockEmbedParams {
	mmEmbed.mutex.RLock()

	argCopy := make([]*OllamaClientInterfaceMockEmbedParams, len(mmEmbed.callArgs))
	copy(argCopy, mmEmbed.callArgs)

	mmEmbed.mutex.RUnlock()

	return argCopy
}

// MinimockEmbedDone returns true if the count of the Embed invocations corresponds
// the number of defined expectations
func (m *OllamaClientInterfaceMock) MinimockEmbedDone() bool {
	if m.EmbedMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.EmbedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.EmbedMock.invocationsDone()
}

// MinimockEmbedInspect logs each unmet expectation
func (m *OllamaClientInterfaceMock) MinimockEmbedInspect() {
	for _, e := range m.EmbedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OllamaClientInterfaceMock.Embed at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterEmbedCounter := mm_atomic.LoadUint64(&m.afterEmbedCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.EmbedMock.defaultExpectation != nil && afterEmbedCounter < 1 {
		if m.EmbedMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to OllamaClientInterfaceMock.Embed at\n%s", m.EmbedMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to OllamaClientInterfaceMock.Embed at\n%s with params: %#v", m.EmbedMock.defaultExpectation.expectationOrigins.origin, *m.EmbedMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcEmbed != nil && afterEmbedCounter < 1 {
		m.t.Errorf("Expected call to OllamaClientInterfaceMock.Embed at\n%s", m.funcEmbedOrigin)
	}

	if !m.EmbedMock.invocationsDone() && afterEmbedCounter > 0 {
		m.t.Errorf("Expected %d calls to OllamaClientInterfaceMock.Embed at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.EmbedMock.expectedInvocations), m.EmbedMock.expectedInvocationsOrigin, afterEmbedCounter)
	}
}

type mOllamaClientInterfaceMockIsAutoPull struct {
	optional           bool
	mock               *OllamaClientInterfaceMock
	defaultExpectation *OllamaClientInterfaceMockIsAutoPullExpectation
	expectations       []*OllamaClientInterfaceMockIsAutoPullExpectation

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// OllamaClientInterfaceMockIsAutoPullExpectation specifies expectation struct of the OllamaClientInterface.IsAutoPull
type OllamaClientInterfaceMockIsAutoPullExpectation struct {
	mock *OllamaClientInterfaceMock

	results      *OllamaClientInterfaceMockIsAutoPullResults
	returnOrigin string
	Counter      uint64
}

// OllamaClientInterfaceMockIsAutoPullResults contains results of the OllamaClientInterface.IsAutoPull
type OllamaClientInterfaceMockIsAutoPullResults struct {
	b1 bool
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmIsAutoPull *mOllamaClientInterfaceMockIsAutoPull) Optional() *mOllamaClientInterfaceMockIsAutoPull {
	mmIsAutoPull.optional = true
	return mmIsAutoPull
}

// Expect sets up expected params for OllamaClientInterface.IsAutoPull
func (mmIsAutoPull *mOllamaClientInterfaceMockIsAutoPull) Expect() *mOllamaClientInterfaceMockIsAutoPull {
	if mmIsAutoPull.mock.funcIsAutoPull != nil {
		mmIsAutoPull.mock.t.Fatalf("OllamaClientInterfaceMock.IsAutoPull mock is already set by Set")
	}

	if mmIsAutoPull.defaultExpectation == nil {
		mmIsAutoPull.defaultExpectation = &OllamaClientInterfaceMockIsAutoPullExpectation{}
	}

	return mmIsAutoPull
}

// Inspect accepts an inspector function that has same arguments as the OllamaClientInterface.IsAutoPull
func (mmIsAutoPull *mOllamaClientInterfaceMockIsAutoPull) Inspect(f func()) *mOllamaClientInterfaceMockIsAutoPull {
	if mmIsAutoPull.mock.inspectFuncIsAutoPull != nil {
		mmIsAutoPull.mock.t.Fatalf("Inspect function is already set for OllamaClientInterfaceMock.IsAutoPull")
	}

	mmIsAutoPull.mock.inspectFuncIsAutoPull = f

	return mmIsAutoPull
}

// Return sets up results that will be returned by OllamaClientInterface.IsAutoPull
func (mmIsAutoPull *mOllamaClientInterfaceMockIsAutoPull) Return(b1 bool) *OllamaClientInterfaceMock {
	if mmIsAutoPull.mock.funcIsAutoPull != nil {
		mmIsAutoPull.mock.t.Fatalf("OllamaClientInterfaceMock.IsAutoPull mock is already set by Set")
	}

	if mmIsAutoPull.defaultExpectation == nil {
		mmIsAutoPull.defaultExpectation = &OllamaClientInterfaceMockIsAutoPullExpectation{mock: mmIsAutoPull.mock}
	}
	mmIsAutoPull.defaultExpectation.results = &OllamaClientInterfaceMockIsAutoPullResults{b1}
	mmIsAutoPull.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmIsAutoPull.mock
}

// Set uses given function f to mock the OllamaClientInterface.IsAutoPull method
func (mmIsAutoPull *mOllamaClientInterfaceMockIsAutoPull) Set(f func() (b1 bool)) *OllamaClientInterfaceMock {
	if mmIsAutoPull.defaultExpectation != nil {
		mmIsAutoPull.mock.t.Fatalf("Default expectation is already set for the OllamaClientInterface.IsAutoPull method")
	}

	if len(mmIsAutoPull.expectations) > 0 {
		mmIsAutoPull.mock.t.Fatalf("Some expectations are already set for the OllamaClientInterface.IsAutoPull method")
	}

	mmIsAutoPull.mock.funcIsAutoPull = f
	mmIsAutoPull.mock.funcIsAutoPullOrigin = minimock.CallerInfo(1)
	return mmIsAutoPull.mock
}

// Times sets number of times OllamaClientInterface.IsAutoPull should be invoked
func (mmIsAutoPull *mOllamaClientInterfaceMockIsAutoPull) Times(n uint64) *mOllamaClientInterfaceMockIsAutoPull {
	if n == 0 {
		mmIsAutoPull.mock.t.Fatalf("Times of OllamaClientInterfaceMock.IsAutoPull mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmIsAutoPull.expectedInvocations, n)
	mmIsAutoPull.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmIsAutoPull
}

func (mmIsAutoPull *mOllamaClientInterfaceMockIsAutoPull) invocationsDone() bool {
	if len(mmIsAutoPull.expectations) == 0 && mmIsAutoPull.defaultExpectation == nil && mmIsAutoPull.mock.funcIsAutoPull == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmIsAutoPull.mock.afterIsAutoPullCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmIsAutoPull.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// IsAutoPull implements OllamaClientInterface
func (mmIsAutoPull *OllamaClientInterfaceMock) IsAutoPull() (b1 bool) {
	mm_atomic.AddUint64(&mmIsAutoPull.beforeIsAutoPullCounter, 1)
	defer mm_atomic.AddUint64(&mmIsAutoPull.afterIsAutoPullCounter, 1)

	mmIsAutoPull.t.Helper()

	if mmIsAutoPull.inspectFuncIsAutoPull != nil {
		mmIsAutoPull.inspectFuncIsAutoPull()
	}

	if mmIsAutoPull.IsAutoPullMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmIsAutoPull.IsAutoPullMock.defaultExpectation.Counter, 1)

		mm_results := mmIsAutoPull.IsAutoPullMock.defaultExpectation.results
		if mm_results == nil {
			mmIsAutoPull.t.Fatal("No results are set for the OllamaClientInterfaceMock.IsAutoPull")
		}
		return (*mm_results).b1
	}
	if mmIsAutoPull.funcIsAutoPull != nil {
		return mmIsAutoPull.funcIsAutoPull()
	}
	mmIsAutoPull.t.Fatalf("Unexpected call to OllamaClientInterfaceMock.IsAutoPull.")
	return
}

// IsAutoPullAfterCounter returns a count of finished OllamaClientInterfaceMock.IsAutoPull invocations
func (mmIsAutoPull *OllamaClientInterfaceMock) IsAutoPullAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsAutoPull.afterIsAutoPullCounter)
}

// IsAutoPullBeforeCounter returns a count of OllamaClientInterfaceMock.IsAutoPull invocations
func (mmIsAutoPull *OllamaClientInterfaceMock) IsAutoPullBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIsAutoPull.beforeIsAutoPullCounter)
}

// MinimockIsAutoPullDone returns true if the count of the IsAutoPull invocations corresponds
// the number of defined expectations
func (m *OllamaClientInterfaceMock) MinimockIsAutoPullDone() bool {
	if m.IsAutoPullMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.IsAutoPullMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.IsAutoPullMock.invocationsDone()
}

// MinimockIsAutoPullInspect logs each unmet expectation
func (m *OllamaClientInterfaceMock) MinimockIsAutoPullInspect() {
	for _, e := range m.IsAutoPullMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to OllamaClientInterfaceMock.IsAutoPull")
		}
	}

	afterIsAutoPullCounter := mm_atomic.LoadUint64(&m.afterIsAutoPullCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.IsAutoPullMock.defaultExpectation != nil && afterIsAutoPullCounter < 1 {
		m.t.Errorf("Expected call to OllamaClientInterfaceMock.IsAutoPull at\n%s", m.IsAutoPullMock.defaultExpectation.returnOrigin)
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIsAutoPull != nil && afterIsAutoPullCounter < 1 {
		m.t.Errorf("Expected call to OllamaClientInterfaceMock.IsAutoPull at\n%s", m.funcIsAutoPullOrigin)
	}

	if !m.IsAutoPullMock.invocationsDone() && afterIsAutoPullCounter > 0 {
		m.t.Errorf("Expected %d calls to OllamaClientInterfaceMock.IsAutoPull at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.IsAutoPullMock.expectedInvocations), m.IsAutoPullMock.expectedInvocationsOrigin, afterIsAutoPullCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *OllamaClientInterfaceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockChatInspect()

			m.MinimockEmbedInspect()

			m.MinimockIsAutoPullInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *OllamaClientInterfaceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *OllamaClientInterfaceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockChatDone() &&
		m.MinimockEmbedDone() &&
		m.MinimockIsAutoPullDone()
}