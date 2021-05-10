const Key = ({ onClick, text, wide, blue, memory}) => {
  return (
    <button
      onClick={onClick}
      className={["key", wide && "wide", blue && "blue", memory && "memory"].join(" ")}
    >
      {text}
    </button>
  );
};

export default Key;
