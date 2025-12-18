import axios from "axios";
import type { queryPosts } from "./types";

const query_url = "http://localhost:4002";

export async function getQueryPosts(): Promise<queryPosts> {
	const req = await axios.get(`${query_url}/posts`);
	return req.data;
}
