import Image from "next/image";

// width and height are the quality of the image and size is the size of the image in html
export default function ProfileImage({ alt, width, height, avatar }) {
  return (
    <div className="overflow-hidden ml-4">
      <Image
        src={avatar}
        width={width}
        height={height}
        className="object-cover h-full w-full rounded-full"
        alt={alt}
        priority={true}
      />
    </div>
  );
}
