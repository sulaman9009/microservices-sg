// post api types
export type post = {
	id: string;
	title: string;
};

export type posts = {
	posts: post[];
};

// comment api types
export type comment = {
	id: string;
	content: string;
};

export type comments = {
	comments: comment[];
};

// query api types
export type queryPost = post & comments;

export type queryPosts = {
	posts: queryPost[];
};
