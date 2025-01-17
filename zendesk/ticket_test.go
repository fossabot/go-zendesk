package zendesk

import (
	"net/http"
	"testing"
)

func TestGetTicket(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticket, err := client.GetTicket(ctx, 2)
	if err != nil {
		t.Fatalf("Failed to get ticket: %s", err)
	}

	expectedID := int64(2)
	if ticket.ID != expectedID {
		t.Fatalf("Returned ticket does not have the expected ID %d. Ticket id is %d", expectedID, ticket.ID)
	}
}

func TestCreateTicket(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "ticket.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticket, err := client.CreateTicket(ctx, Ticket{
		Subject: "nyanyanyanya",
		Comment: TicketComment{
			Body: "(●ↀ ω ↀ )",
		},
	})
	if err != nil {
		t.Fatalf("Failed to create ticket: %s", err)
	}

	expectedID := int64(4)
	if ticket.ID != expectedID {
		t.Fatalf("Returned ticket does not have the expected ID %d. Ticket id is %d", expectedID, ticket.ID)
	}
}
