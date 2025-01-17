package password

// Built-time checks that the generators implement the interface.
var _ PasswordGenerator = (*mockGenerator)(nil)

type mockGenerator struct {
	result string
	err    error
}

// NewMockGenerator creates a new generator that satisfies the PasswordGenerator
// interface. If an error is provided, the error is returned. If a result if
// provided, the result is always returned, regardless of what parameters are
// passed into the Generate or MustGenerate methods.
//
// This function is most useful for tests where you want to have predicable
// results for a transitive resource that depends on go-password.
func NewMockGenerator(result string, err error) *mockGenerator {
	return &mockGenerator{
		result: result,
		err:    err,
	}
}

// Generate returns the mocked result or error.
func (g *mockGenerator) Generate(int, int, int, bool, bool, string) (string, error) {
	if g.err != nil {
		return "", g.err
	}
	return g.result, nil
}

// MustGenerate returns the mocked result or panics if an error was given.
func (g *mockGenerator) MustGenerate(int, int, int, bool, bool, string) string {
	if g.err != nil {
		panic(g.err)
	}
	return g.result
}
