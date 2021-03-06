// This package was auto generated.
// DO NOT EDIT BY YOUR HAND!

package errmsg

import "fmt"

// StructNameIsNilErr returns the error.
func StructNameIsNilErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-1] struct name must not be empty, but it gets empty (caused at %s)`, caller)
}

// StructFieldNameIsEmptyErr returns the error.
func StructFieldNameIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-2] field name must not be empty, but it gets empty (caused at %s)`, caller)
}

// StructFieldTypeIsEmptyErr returns the error.
func StructFieldTypeIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-3] field type must not be empty, but it gets empty (caused at %s)`, caller)
}

// FuncParameterNameIsEmptyErr returns the error.
func FuncParameterNameIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-4] func parameter name must not be empty, but it gets empty (caused at %s)`, caller)
}

// LastFuncParameterTypeIsEmptyErr returns the error.
func LastFuncParameterTypeIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-5] the last func parameter type must not be empty, but it gets empty (caused at %s)`, caller)
}

// FuncNameIsEmptyError returns the error.
func FuncNameIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-6] name of func must not be empty, but it gets empty (caused at %s)`, caller)
}

// InterfaceNameIsEmptyError returns the error.
func InterfaceNameIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-7] name of interface must not be empty, but it gets empty (caused at %s)`, caller)
}

// FuncReceiverNameIsEmptyError returns the error.
func FuncReceiverNameIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-8] name of func receiver must not be empty, but it gets empty (caused at %s)`, caller)
}

// FuncReceiverTypeIsEmptyError returns the error.
func FuncReceiverTypeIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-9] type of func receiver must not be empty, but it gets empty (caused at %s)`, caller)
}

// FuncSignatureIsNilError returns the error.
func FuncSignatureIsNilError(caller string) error {
	return fmt.Errorf(`[GOWRTR-10] func signature must not be nil, bit it gets nil (caused at %s)`, caller)
}

// AnonymousFuncSignatureIsNilError returns the error.
func AnonymousFuncSignatureIsNilError(caller string) error {
	return fmt.Errorf(`[GOWRTR-11] anonymous func signature must not be nil, bit it gets nil (caused at %s)`, caller)
}

// FuncInvocationParameterIsEmptyError returns the error.
func FuncInvocationParameterIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-12] a parameter of function invocation must not be nil, but it gets nil (caused at %s)`, caller)
}

// CodeFormatterError returns the error.
func CodeFormatterError(cmd string, msg string, err error) error {
	return fmt.Errorf(`[GOWRTR-13] code formatter raises error: command="%s", err="%s", msg="%s"`, cmd, msg, err)
}

// CaseConditionIsEmptyError returns the error.
func CaseConditionIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-14] condition of case must not be empty, but it gets empty (caused at %s)`, caller)
}

// IfConditionIsEmptyError returns the error.
func IfConditionIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-15] condition of if must not be empty, but it gets empty (caused at %s)`, caller)
}

// UnnamedReturnTypeAppearsAfterNamedReturnTypeError returns the error.
func UnnamedReturnTypeAppearsAfterNamedReturnTypeError(caller string) error {
	return fmt.Errorf(`[GOWRTR-16] unnamed return type appears after named return type (caused at %s)`, caller)
}

// ValueOfCompositeLiteralIsEmptyError returns the error.
func ValueOfCompositeLiteralIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-17] a value of composite literal must not be empty, but it gets empty (caused at %s)`, caller)
}

// ErrsList returns the list of errors.
func ErrsList() []string {
	return []string{"[GOWRTR-1] struct name must not be empty, but it gets empty (caused at %s)", "[GOWRTR-2] field name must not be empty, but it gets empty (caused at %s)", "[GOWRTR-3] field type must not be empty, but it gets empty (caused at %s)", "[GOWRTR-4] func parameter name must not be empty, but it gets empty (caused at %s)", "[GOWRTR-5] the last func parameter type must not be empty, but it gets empty (caused at %s)", "[GOWRTR-6] name of func must not be empty, but it gets empty (caused at %s)", "[GOWRTR-7] name of interface must not be empty, but it gets empty (caused at %s)", "[GOWRTR-8] name of func receiver must not be empty, but it gets empty (caused at %s)", "[GOWRTR-9] type of func receiver must not be empty, but it gets empty (caused at %s)", "[GOWRTR-10] func signature must not be nil, bit it gets nil (caused at %s)", "[GOWRTR-11] anonymous func signature must not be nil, bit it gets nil (caused at %s)", "[GOWRTR-12] a parameter of function invocation must not be nil, but it gets nil (caused at %s)", "[GOWRTR-13] code formatter raises error: command=\"%s\", err=\"%s\", msg=\"%s\"", "[GOWRTR-14] condition of case must not be empty, but it gets empty (caused at %s)", "[GOWRTR-15] condition of if must not be empty, but it gets empty (caused at %s)", "[GOWRTR-16] unnamed return type appears after named return type (caused at %s)", "[GOWRTR-17] a value of composite literal must not be empty, but it gets empty (caused at %s)"}
}
