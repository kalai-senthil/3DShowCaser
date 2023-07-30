export type User = {
	email: string;
	id: string;
	works: Work[];
};
export type Work = {
	id: string;
	name: string;
	uploadedAt: string;
	file?: string;
};
