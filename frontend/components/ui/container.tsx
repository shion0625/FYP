interface ContainerProps {
  children: React.ReactNode;
}

const Container: React.FC<ContainerProps> = ({ children }) => (
  <div className="mx-auto max-w-7xl">{children}</div>
);

export default Container;
