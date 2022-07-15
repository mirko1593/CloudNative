FROM ruby:2.3.1
RUN mkdir /lobsters
WORKDIR /lobsters

ADD Gemfile /lobsters/Gemfile
ADD Gemfile.lock /lobsters/Gemfile.lock
RUN apt-get update && apt-get install nodejs -y
RUN bundle install
ADD . /lobsters
ENTRYPOINT ["/lobsters/bin/rails", "server"]
