package context

import "context"

type contextKey string

const (
	RequestIDKey   contextKey = "request_id"
	UserIDKey      contextKey = "user_id"
	InstitutionKey contextKey = "institution_id"
	TraceIDKey     contextKey = "trace_id"
)

//
// REQUEST ID
//

func WithRequestID(
	ctx context.Context,
	reqID string,
) context.Context {
	return context.WithValue(
		ctx,
		RequestIDKey,
		reqID,
	)
}

func RequestID(
	ctx context.Context,
) string {
	val, ok := ctx.Value(
		RequestIDKey,
	).(string)

	if !ok {
		return ""
	}

	return val
}

//
// USER ID
//

func WithUserID(
	ctx context.Context,
	userID string,
) context.Context {
	return context.WithValue(
		ctx,
		UserIDKey,
		userID,
	)
}

func UserID(
	ctx context.Context,
) string {
	val, ok := ctx.Value(
		UserIDKey,
	).(string)

	if !ok {
		return ""
	}

	return val
}

//
// INSTITUTION ID
//

func WithInstitutionID(
	ctx context.Context,
	institutionID string,
) context.Context {
	return context.WithValue(
		ctx,
		InstitutionKey,
		institutionID,
	)
}

func InstitutionID(
	ctx context.Context,
) string {
	val, ok := ctx.Value(
		InstitutionKey,
	).(string)

	if !ok {
		return ""
	}

	return val
}

//
// TRACE ID
//

func WithTraceID(
	ctx context.Context,
	traceID string,
) context.Context {
	return context.WithValue(
		ctx,
		TraceIDKey,
		traceID,
	)
}

func TraceID(
	ctx context.Context,
) string {
	val, ok := ctx.Value(
		TraceIDKey,
	).(string)

	if !ok {
		return ""
	}

	return val
}
