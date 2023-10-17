package dal

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ravibhagw/xbs_adv_stats_job/models"
)

// Client represents a MongoDB client for storing data.
type Client struct {
	client *mongo.Client
}

// NewClient creates a new Client instance.
func NewClient(connectionString string) (*Client, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}

// Initialize opens the MongoDB connection.
func (c *Client) Initialize() error {
	err := c.client.Connect(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Close closes the MongoDB connection.
func (c *Client) Close() {
	if err := c.client.Disconnect(context.Background()); err != nil {
		log.Println("Error while closing MongoDB connection:", err)
	}
}

func (c *Client) SaveMatches(matches []models.Match) {
	// Implement the logic to save matches here
}

// SaveClubData saves club data to the MongoDB collection.
func (c *Client) SaveClubData(clubs []models.Team) {
	// Implement the logic to save club data here
}

// SavePlayerData saves player data to the MongoDB collection.
func (c *Client) SavePlayerData(players []models.Player) {
	// Implement the logic to save player data here
}
