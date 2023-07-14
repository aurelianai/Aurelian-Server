export type Chat = {
	ID: number;
	Title: string;
	CreatedAt: string;
};

export type Message = {
	Role: 'USER' | 'MODEL';
	Content: string;
};

export type User = {
	Email: string,
	FirstName: string,
	LastName: string
}