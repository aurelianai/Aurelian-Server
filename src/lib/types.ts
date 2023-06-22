export type Chat = {
    id: number;
    name: string;
};

export type Message = {
	role: 'USER' | 'MODEL';
	content: string;
};

export type TextGenModel = {
	id: number;
	name: string;
	description: string;
};