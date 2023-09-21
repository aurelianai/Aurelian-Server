export type Database = {
	user: UserTable;
	key: KeyTable;
	session: SessionTable;
};

type UserTable = {
   id: string;
};

type KeyTable = {
	id: string;
	user_id: string;
	hashed_password: string | null;
};

type SessionTable = {
	id: string;
	user_id: string;
	active_expires: number;
	idle_expires: number;
};