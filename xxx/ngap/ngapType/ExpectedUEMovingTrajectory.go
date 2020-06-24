// Created By HaoDHH-245789 VHT2020
package ngapType

/* Sequence of = 35, FULL Name = struct ExpectedUEMovingTrajectory */
/* ExpectedUEMovingTrajectoryItem */
type ExpectedUEMovingTrajectory struct {
	List []ExpectedUEMovingTrajectoryItem `vht:"valueExt,sizeLB:1,sizeUB:16"`
}
