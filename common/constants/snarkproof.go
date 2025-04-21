package constants

type CircuitType string
const (
	APP_CIRCUIT  CircuitType = "app"
	CLIENT_PAYLOAD  CircuitType = "client_payload"
)
const (
	VERFICATION_KEY_JSON_FILE = "verification_key.json"
	PROVING_KEY_FILE = "proving_key.pk"
	R1CS_FILE = "circuit.r1cs"
)