'use client';
import { SkeletonTheme } from 'react-loading-skeleton';
import 'react-loading-skeleton/dist/skeleton.css';

interface loadingSkeletonProps {
  baseColor?: string;
  highlightColor?: string;
  children: React.ReactNode;
}
const LoadingSkeleton: React.FC<loadingSkeletonProps> = ({
  baseColor = '#ebebeb',
  highlightColor = '#f5f5f5',
  children,
}) => {
  return (
    <SkeletonTheme baseColor={baseColor} highlightColor={highlightColor}>
      {children}
    </SkeletonTheme>
  );
};

export default LoadingSkeleton;
