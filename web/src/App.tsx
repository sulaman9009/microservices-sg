import CreatePostsForm from "./components/posts/create-posts-form";
import PostList from "./components/posts/post-list";

function App() {
	return (
		<div>
			<h1 className="mb-2 ml-2">Create A Post</h1>
			<CreatePostsForm />
			<hr />
			<h2 className="mb-2 ml-2">Posts</h2>
			<PostList />
		</div>
	);
}

export default App;
