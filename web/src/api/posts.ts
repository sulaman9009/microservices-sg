import axios from "axios";
import type { post, posts } from "./types";

const posts_url = "http://localhost:4000";

export async function createPost(title: string): Promise<post> {
	const req = await axios.post(`${posts_url}/post`, {
		title: title,
	});
	return req.data;
}

export async function getPosts(): Promise<posts> {
	const req = await axios.get(`${posts_url}/posts`);
	return req.data;
}
