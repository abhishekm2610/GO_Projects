class UndergroundSystem:
    def __init__(self):
        self.OngoingTrips = {}
        self.AvgTime = {}

    def checkIn(self, id: int, stationName: str, t: int) -> None:
        self.OngoingTrips[id] = (stationName, t)
        return

    def checkOut(self, id: int, stationName: str, t: int) -> None:
        totalTime = t - self.OngoingTrips[id][1]
        startStation = self.OngoingTrips[id][0]
        if (startStation, stationName) not in self.AvgTime:
            self.AvgTime[(startStation, stationName)] = [0, 0]
        self.AvgTime[(startStation, stationName)][0] += totalTime
        self.AvgTime[(startStation, stationName)][1] += 1

        del self.OngoingTrips[id]

        return

    def getAverageTime(self, startStation: str, endStation: str) -> float:
        return (
            self.AvgTime[(startStation, endStation)][0]
            / self.AvgTime[(startStation, endStation)][1]
        )


# Your UndergroundSystem object will be instantiated and called as such:
# obj = UndergroundSystem()
# obj.checkIn(id,stationName,t)
# obj.checkOut(id,stationName,t)
# param_3 = obj.getAverageTime(startStation,endStation)
