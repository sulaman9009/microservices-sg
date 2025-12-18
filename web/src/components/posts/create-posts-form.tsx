import { useActionState } from "react";
import { createPost } from "../../api/posts";

function CreatePostsForm() {
	const submitForm = async (_: unknown, data: FormData) => {
		const title = data.get("title")?.toString();
		if (!title) {
			console.error("title not set");
			return { result: "failure" };
		}
		try {
			await createPost(title);
		} catch (error) {
			console.error("failed to create post:", error);
			return { result: "failure" };
		}
		return { result: "success" };
	};

	const [_, action, isLoading] = useActionState(submitForm, undefined);

	return (
		<div className="w-[50%] p-3 mx-auto">
			<form action={action}>
				<div className="flex flex-col">
					<label htmlFor="title" className="text-zinc-800 text-md">
						Title
					</label>
					<input
						type="text"
						name="title"
						id="title"
						className="border-zinc-400 rounded-md border p-1"
					/>
				</div>
				<button
					type="submit"
					className="bg-sky-400 text-zinc-800 py-1 px-2 mt-3 rounded-md hover:bg-blue-500 transition-colors ease-in-out cursor-pointer"
				>
					{isLoading ? "submitting..." : "submit"}
				</button>
			</form>
		</div>
	);
}

export default CreatePostsForm;
