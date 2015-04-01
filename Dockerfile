FROM scratch
ADD gohello gohello
EXPOSE 80 1000 2000 3000 4000 5000
ENTRYPOINT ["/gohello"]
CMD ["80", "1000", "2000", "3000", "4000", "5000"]

