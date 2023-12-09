'use client';

import { FaCcVisa, FaCcMastercard, FaCcAmex, FaCcDiscover, FaCcDinersClub } from 'react-icons/fa';

interface CardIconProps {
  cardCompany: string;
}

const CardIcon: React.FC<CardIconProps> = ({ cardCompany }) => {
  switch (cardCompany) {
    case 'Visa':
      return <FaCcVisa/>;
    case 'MasterCard':
      return <FaCcMastercard/>;
    case 'American Express':
      return <FaCcAmex/>;
    case 'Discover':
      return <FaCcDiscover/>;
    case 'Diners Club':
      return <FaCcDinersClub/>;
    default:
      return <></>; // default icon
  }
};

export default CardIcon;
