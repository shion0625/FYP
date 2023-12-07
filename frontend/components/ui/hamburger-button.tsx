import { Button } from 'flowbite-react';
import { GiHamburgerMenu } from 'react-icons/gi';
import { RxCross1 } from 'react-icons/rx';

interface HamburgerProps {
  isOpen: boolean;
  onOpen: () => void;
  onClose: () => void;
}
const HamburgerMenu: React.FC<HamburgerProps> = ({ isOpen, onOpen, onClose }) => (
  <>
    {isOpen ? (
      <Button outline onClick={onClose} className="mr-8">
        <RxCross1 />
      </Button>
    ) : (
      <Button outline onClick={onOpen} className="mr-8">
        <GiHamburgerMenu />
      </Button>
    )}
  </>
);

export default HamburgerMenu;
