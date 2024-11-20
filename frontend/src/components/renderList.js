import Link from "next/link";

const RenderList = ({ items, showMoreState, setShowMoreState, type }) => {
  return (
    <>
      <ul className="list-disc pl-5 marker:text-txtColor">
        {(showMoreState ? items : items.slice(0, 4)).map((item, index) => (
          <li key={index}>
            <Link
              href={`/${
                type === "event"
                  ? "events"
                  : type === "group"
                  ? "group"
                  : "profile"
              }/${item}`}
              className="text-txtColor hover:underline"
            >
              {item}
            </Link>
          </li>
        ))}
      </ul>
      {items.length > 4 && showMore(showMoreState, setShowMoreState)}
    </>
  );
};

const showMore = (state, setState) => {
  return (
    <button
      onClick={() => setState(!state)}
      className="hover:underline flex items-center"
    >
      {state ? "Show Less" : "Show More"}
      <span className={`ml-1 transform ${state ? "rotate-180" : ""}`}>▼</span>
    </button>
  );
};

export default RenderList;
