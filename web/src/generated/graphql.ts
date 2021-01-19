import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Task = {
  __typename?: 'Task';
  createdAt: Scalars['String'];
  createdBy: User;
  id: Scalars['ID'];
  text: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  loggedInUser?: Maybe<User>;
  tasks: Array<Task>;
};

export type User = {
  __typename?: 'User';
  avatar: Scalars['String'];
  id: Scalars['ID'];
  name: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createTask: Task;
};


export type MutationCreateTaskArgs = {
  input: CreateTask;
};

export type CreateTask = {
  text: Scalars['String'];
};

export type LoggedInUserQueryVariables = Exact<{ [key: string]: never; }>;


export type LoggedInUserQuery = (
  { __typename?: 'Query' }
  & { loggedInUser?: Maybe<(
    { __typename?: 'User' }
    & Pick<User, 'avatar' | 'id' | 'name'>
  )> }
);


export const LoggedInUserDocument = gql`
    query LoggedInUser {
  loggedInUser {
    avatar
    id
    name
  }
}
    `;

/**
 * __useLoggedInUserQuery__
 *
 * To run a query within a React component, call `useLoggedInUserQuery` and pass it any options that fit your needs.
 * When your component renders, `useLoggedInUserQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useLoggedInUserQuery({
 *   variables: {
 *   },
 * });
 */
export function useLoggedInUserQuery(baseOptions?: Apollo.QueryHookOptions<LoggedInUserQuery, LoggedInUserQueryVariables>) {
        return Apollo.useQuery<LoggedInUserQuery, LoggedInUserQueryVariables>(LoggedInUserDocument, baseOptions);
      }
export function useLoggedInUserLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<LoggedInUserQuery, LoggedInUserQueryVariables>) {
          return Apollo.useLazyQuery<LoggedInUserQuery, LoggedInUserQueryVariables>(LoggedInUserDocument, baseOptions);
        }
export type LoggedInUserQueryHookResult = ReturnType<typeof useLoggedInUserQuery>;
export type LoggedInUserLazyQueryHookResult = ReturnType<typeof useLoggedInUserLazyQuery>;
export type LoggedInUserQueryResult = Apollo.QueryResult<LoggedInUserQuery, LoggedInUserQueryVariables>;

      export interface PossibleTypesResultData {
        possibleTypes: {
          [key: string]: string[]
        }
      }
      const result: PossibleTypesResultData = {
  "possibleTypes": {}
};
      export default result;
    