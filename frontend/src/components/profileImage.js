import Image from "next/image";

// width and height are the quality of the image and size is the size of the image in html
export default function ProfileImage({ alt, width, height, size, userId }) {
  // TODO: Fetch user profile image from images using userId
  const username = "userName"; // placeholder
  const src = "/image/avatar/" + username;
  return (
    <div className="overflow-hidden">
      <div style={{ width: size, height: size }}>
        <Image
          src={src}
          width={width}
          height={height}
          className="object-cover h-full w-full rounded-full"
          alt={alt}
          priority={true}
        />
      </div>
    </div>
  );
}
