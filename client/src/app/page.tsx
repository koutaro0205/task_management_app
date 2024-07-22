'use client';

import type React from 'react';

import { Heading } from '@/components/ui/atoms/Heading';
import { TodoList } from '@/model/todo/components/TodoList';

const Home: React.FC = () => {
  return (
    <div>
      <Heading as="h2">Hello World !!</Heading>
      <TodoList />
    </div>
  );
};

export default Home;
