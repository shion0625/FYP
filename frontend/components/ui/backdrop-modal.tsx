'use client';

import { useState } from 'react';
import { Button, Modal } from 'flowbite-react';

interface BackdropModalProps {
  headerText: string;
  buttonText: string;
  buttonClassName: string;
  children: React.ReactNode;
}
const BackdropModal: React.FC<BackdropModalProps> = ({
  children,
  buttonClassName,
  buttonText,
  headerText,
}) => {
  const [openModal, setOpenModal] = useState(false);

  return (
    <>
      <Button color="white" className={buttonClassName} onClick={() => setOpenModal(true)}>
        {buttonText}
      </Button>
      <Modal dismissible show={openModal} onClose={() => setOpenModal(false)}>
        <Modal.Header>{headerText}</Modal.Header>
        <Modal.Body>{children}</Modal.Body>
      </Modal>
    </>
  );
};

export default BackdropModal;
