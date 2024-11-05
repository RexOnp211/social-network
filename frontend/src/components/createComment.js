import FetchFromBackend from "@/lib/fetch";

export default function CreateComment({ postId }) {
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const form = new FormData(e.target);
      const response = await FetchFromBackend(`/post/${postId}`, {
        method: "POST",
        credentials: "include",
        headers: {},
        body: form,
      });
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
    } catch (error) {
      console.error("Error submitting the form:", error);
    }
  };
  return (
    <div className="flex flex-col w-full">
      <form
        id="CreateComment"
        className="flex flex-col justify-center border border-gray-300 p-6 rounded-lg shadow-md bg-white"
        onSubmit={handleSubmit}
      >
        <label
          htmlFor="commentBody"
          className="mb-2 text-lg font-semibold text-gray-700"
        >
          {" "}
          Comment
        </label>
        <textarea
          id="commentBody"
          name="commentBody"
          placeholder="Enter comment body"
          className="mb-4 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-accent focus:bg-primary"
        ></textarea>
        <label
          htmlFor="image"
          className="mb-2 text-lg font-semibold text-gray-700"
        >
          Upload Image
        </label>
        <input
          type="file"
          id="image"
          name="image"
          className="mb-4 w-full p-2 border border-gray-300 rounded-md focus:bg-primary"
        />
        <button
          type="submit"
          className="bg-accent w-full text-white rounded-lg p-2 transition-colors hover:bg-accentDark"
        >
          Create Comment
        </button>
      </form>
    </div>
  );
}
