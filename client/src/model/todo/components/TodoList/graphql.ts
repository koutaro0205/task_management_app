import { graphql } from '@/gql';

export const GET_TODOS = graphql(`
  query GetTodos {
    todos {
      id
      title
      done
      createdAt
      user {
        id
      }
    }
  }
`);
