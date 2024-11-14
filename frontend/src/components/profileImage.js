import Image from "next/image";

// width and height are the quality of the image and size is the size of the image in html
export default function ProfileImage({
  alt,
  width,
  height,
  size,
  avatar,
  className,
}) {
  return (
    <div className="overflow-hidden">
      <Image
        src={avatar}
        width={width}
        height={height}
        alt={alt}
        className={className}
        priority={true}
      />
    </div>
  );
}
